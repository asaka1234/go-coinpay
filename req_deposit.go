package go_coinpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-coinpay/utils"
	goquery "github.com/google/go-querystring/query"
	jsoniter "github.com/json-iterator/go"
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
	bodyForm.Add("success_url", cli.Params.DepositFeBackUrl)
	bodyForm.Add("cancel_url", cli.Params.DepositFeBackUrl)

	//计算sign (要放在Head里)
	payload := bodyForm.Encode()

	mac := hmac.New(sha512.New, []byte(cli.Params.PrivateKey)) //
	mac.Write([]byte(payload))
	hmac := fmt.Sprintf("%x", mac.Sum(nil))

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

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2, payload))
	cli.logger.Infof("PSPResty#coinpay#deposit->%+v", string(restLog))

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
	}, fmt.Errorf("%s", result.Error)
}
