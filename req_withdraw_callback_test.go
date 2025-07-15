package go_coinpay

//--------------------------------------------

/*
func TestWithdrawCallback(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &CoinPayInitParams{MerchantID, PrivateKey, PublicKey, IPNSecret, EndPoint, DepositBackUrl, WithdrawBackUrl})

	//发请求
	resp, err := cli.Deposit(GenWithdrawCallbackRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawCallbackRequestDemo() string {

	return CoinPayDepositReq{
		Currency1:  "USD",
		Currency2:  "USDT.TRC20",        //"USDT.TRC20",
		BuyerEmail: "1609032335@qq.com", //outNo
		Amount:     "1",
		Invoice:    "1232131", //商户订单号
		Custom:     "89071",   //uid
	}
}

*/
