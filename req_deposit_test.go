package go_coinpay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

//--------------------------------------------

func TestDeposit(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, CoinPayInitParams{MerchantID, PrivateKey, PublicKey, IPNSecret, EndPoint, DepositBackUrl, WithdrawBackUrl})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() CoinPayDepositReq {

	return CoinPayDepositReq{
		Currency1:  "USD",
		Currency2:  "LTCT",              //"USDT.TRC20",
		BuyerEmail: "1609032335@qq.com", //outNo
		Amount:     "1",
		Invoice:    "123213", //商户订单号
		Custom:     "8907",   //uid
	}
}
