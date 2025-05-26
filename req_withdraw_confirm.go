package go_buy365

import (
	"crypto/tls"
	"github.com/asaka1234/go-buy365/utils"
	"github.com/mitchellh/mapstructure"
)

// withdraw确认
func (cli *Client) WithdrawConfirm(req Buy365WithdrawConfirmReq) (*Buy365WithdrawConfirmResponse, error) {

	rawURL := cli.WithdrawConfirmURL

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["sys_no"] = cli.MerchantID

	//签名
	signStr := utils.SignDeposit(params, cli.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result Buy365WithdrawConfirmResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
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
