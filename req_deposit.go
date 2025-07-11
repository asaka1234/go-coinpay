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

// https://www.coinpayments.net/apidoc-create-transaction
func (cli *Client) Deposit(req CoinPayDepositReq) (*CoinPayDepositResponse, error) {

	rawURL := cli.Params.EndPoint

	//构造请求body
	bodyForm, err := goquery.Values(req)
	if err != nil {
		log.Fatal(err)
	}
	//添加公共参数
	bodyForm.Add("key", cli.Params.PublicKey)
	bodyForm.Add("version", "1") //FIXED
	bodyForm.Add("cmd", "create_transaction")
	bodyForm.Add("format", "json") //FIXED
	bodyForm.Add("ipn_url", cli.Params.DepositBackUrl)

	//计算sign (要放在Head里)
	payload := bodyForm.Encode()

	fmt.Printf("===>payload:%s\n", payload)

        mac := hmac.New(sha512.New, []byte(cli.Params.PrivateKey))  //
	mac.Write([]byte(payload))
	hmac := fmt.Sprintf("%x", mac.Sum(nil))

	fmt.Printf("===>sign:%s\n", hmac)

	//返回值会放到这里
	var result CoinPayDepositCommonResponse

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders(hmac)).
	        SetDebug(cli.debugMode).
		SetFormDataFromValues(bodyForm).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
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
		var resp3 CoinPayDepositResponse
		if err := mapstructure.Decode(data, &resp3); err != nil {
			return nil, err
		}

		return &resp3, nil
	}

	return &CoinPayDepositResponse{
		Error: result.Error,
	}, fmt.Errorf(result.Error)
}
