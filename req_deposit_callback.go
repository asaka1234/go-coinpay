package go_coinpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

// body是http-body的整体内容.
func (cli *Client) DepositCallback(body string, hmacHeader string, processor func(CoinPayDepositBackReq) error) error {
	//1. 验证签名
	//计算HMAC签名
	mac := hmac.New(sha512.New, []byte(cli.Params.IPNSecret))
	mac.Write([]byte(body))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	//验证签名
	if !hmac.Equal([]byte(expectedMAC), []byte(hmacHeader)) {
		return fmt.Errorf("HMAC signature does not match")
	}

	//-------------------

	//step-1
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return err
	}

	//step-2
	var resp3 CoinPayDepositBackReq
	if err := mapstructure.Decode(data, &resp3); err != nil {
		return err
	}

	//开始处理
	return processor(resp3)
}
