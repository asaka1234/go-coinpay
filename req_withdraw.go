package go_coinpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"crypto/tls"
	"encoding/json"
	"fmt"
	goquery "github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
	"log"
	"strings"
)

// withdraw
func (cli *Client) Withdraw(req CoinPayWithdrawalRequest) (*CoinPayWithdrawalResponse, error) {

	rawURL := cli.Params.EndPoint

	//构造请求body
	bodyForm, err := goquery.Values(req)
	if err != nil {
		log.Fatal(err)
	}
	//添加公共参数
	bodyForm.Add("key", cli.Params.PublicKey)
	bodyForm.Add("version", "1") //FIXED
	bodyForm.Add("cmd", "create_withdrawal")
	bodyForm.Add("format", "json") //FIXED
	bodyForm.Add("ipn_url", cli.Params.WithdrawBackUrl)

	//计算sign (要放在Head里)
	payload := bodyForm.Encode()

	fmt.Printf("===>payload:%s\n", payload)

	mac := hmac.New(sha512.New, []byte(cli.Params.PrivateKey))
	mac.Write([]byte(payload))
	hmac := fmt.Sprintf("%x", mac.Sum(nil))

	fmt.Printf("===>sign:%s\n", hmac)

	//返回值会放到这里
	var result CoinPayWithdrawalCommonResponse

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders(hmac)).
	        SetDebug(cli.debugMode).
	        SetLogger(cli.logger).
		SetFormDataFromValues(bodyForm).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	responseStr := string(resp2.Body())
	log.Printf("CoinPayService#deposit#rsp: %s", responseStr)

	if strings.ToLower(result.Error) == "ok" {
		//说明成功了

		//step-1
		var data map[string]interface{}
		if err := json.Unmarshal(resp2.Body(), &data); err != nil {
			return nil, err
		}

		//step-2
		var resp3 CoinPayWithdrawalResponse
		if err := mapstructure.Decode(data, &resp3); err != nil {
			return nil, err
		}

		return &resp3, nil
	}

	return &CoinPayWithdrawalResponse{
		Error: result.Error,
	}, fmt.Errorf(result.Error)
}
