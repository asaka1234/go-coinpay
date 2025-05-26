package go_buy365

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, MERCHANT_ID, ACCESS_KEY, BACK_KEY, DEPOSIT_URL, WITHDRAW_URL, WITHDRAW_CONFIRM_URL, ORDERLIST_URL)

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() Buy365DepositReq {
	return Buy365DepositReq{
		OrderId:     "3234", //商户uid
		UserId:      "30779639363",
		OrderIp:     "18.29.120.32",
		OrderAmount: "60000.00",
		PayUserName: "你好", //商户订单号
	}
}
