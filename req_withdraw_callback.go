package go_coinpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// body是http-body的整体内容.
func (cli *Client) WithdrawCallback(body string, hmacHeader string, req CoinPayWithdrawalBackReq, processor func(req CoinPayWithdrawalBackReq) error) error {
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

	//开始处理
	return processor(req)
}
