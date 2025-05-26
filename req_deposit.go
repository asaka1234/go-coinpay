package go_buy365

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-buy365/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// pre-order
func (cli *Client) Deposit(req Buy365DepositReq) (*Buy365DepositResponse, error) {

	rawURL := cli.DepositURL

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["sys_no"] = cli.MerchantID
	params["order_time"] = time.Now().Format("2006-01-02 15:04:05")

	//签名
	signStr := utils.SignDeposit(params, cli.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result Buy365DepositCommonResponse

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	//------------------------------------------------------
	if result.Code == 111 && result.Status == "success" {
		//说明成功

		//step-1
		var data map[string]interface{}
		if err := json.Unmarshal(resp2.Body(), &data); err != nil {
			return nil, err
		}

		//step-2
		var resp3 Buy365DepositResponse
		if err := mapstructure.Decode(data, &resp3); err != nil {
			return nil, err
		}

		return &resp3, nil
	}

	return &Buy365DepositResponse{
		Code:   result.Code,
		Status: result.Status,
		Msg:    result.Msg,
	}, fmt.Errorf("result is failed")
}
