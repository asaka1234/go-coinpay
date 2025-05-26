package go_coinpay

import (
	"crypto/tls"
	"encoding/json"
	"github.com/asaka1234/go-coinpay/utils"
)

// withdraw
func (cli *Client) Withdraw(req CoinPayWithdrawReq) (*CoinPayWithdrawResponse, error) {

	rawURL := cli.WithdrawURL

	jsonData, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["data"] = string(jsonData)
	params["sys_no"] = cli.MerchantID

	//签名
	signStr := utils.SignWithdraw(params, cli.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result CoinPayWithdrawResponse

	_, err = cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}
