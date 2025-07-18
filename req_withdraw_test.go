package go_coinpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"testing"
)

//--------------------------------------------

func TestWithdraw(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &CoinPayInitParams{MerchantID, PrivateKey, PublicKey, IPNSecret, EndPoint, DepositBackUrl, WithdrawBackUrl, DepositFeBackUrl})
	cli.SetDebugModel(true)
	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() CoinPayWithdrawalRequest {

	return CoinPayWithdrawalRequest{
		Currency:    "USDT.TRC20",
		Address:     "la st",
		Amount:      "0.5",
		AutoConfirm: 1,       //商户订单号
		Note:        "89071", //uid
	}
}

func TestWithdraw2(t *testing.T) {
	body := "amount1=1&amount2=6.60496&buyer_name=CoinPayments+API&currency1=USD&currency2=USDT.TRC20&custom=89071&email=1609032335%40qq.com&fee=5.61951&invoice=1232131114&ipn_id=139bd7a8a46c2d8009ad3f95329c7104&ipn_mode=hmac&ipn_type=api&ipn_version=1.0&merchant=57bb15a3e423b4e2d1a1d76477240d5c&received_amount=0&received_confirms=0&status=0&status_text=Waiting+for+buyer+funds...&txn_id=CPJG3K24VV98VGFT35YRYDYT1V"
	//"amount1=1&amount2=6.58186&buyer_name=CoinPayments+API&currency1=USD&currency2=USDT.TRC20&custom=89071&email=1609032335%40qq.com&fee=5.59735&invoice=1232131113&ipn_id=dec9bc2f7925637fd76ed789403d9aa8&ipn_mode=hmac&ipn_type=api&ipn_version=1.0&merchant=57bb15a3e423b4e2d1a1d76477240d5c&received_amount=0&received_confirms=0&status=0&status_text=Waiting+for+buyer+funds...&txn_id=CPJG1LFU6IYKKBQJNG41D8J1UJ"

	mac := hmac.New(sha512.New, []byte("CPTint2022@F"))
	mac.Write([]byte(body))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	fmt.Printf("==>%s\n", expectedMAC)

}
