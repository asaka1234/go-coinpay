package go_coinpay

// ----------pre order-------------------------

// https://www.coinpayments.net/apidoc-create-transaction
type CoinPayDepositReq struct {
	//比如我们想让用户充usd, 但是希望收到的是btc.  则currency1=USD, currency2=BTC

	// Required fields
	Amount     string `json:"amount" mapstructure:"amount" url:"amount"`                        // The amount of the payment in the original currency - Yes
	Currency1  string `json:"currency1" mapstructure:"currency1" url:"currency1"`               // The original currency of the payment.
	Currency2  string `json:"currency2" mapstructure:"currency2" url:"currency2"`               // The currency the buyer will be sending - Yes
	BuyerEmail string `json:"buyer_email" mapstructure:"buyer_email" url:"buyer_email"`         // Buyer's email address - Yes
	Invoice    string `json:"invoice,omitempty" mapstructure:"invoice,omitempty" url:"invoice"` //option, 商户订单号
	Custom     string `json:"custom,omitempty" mapstructure:"custom,omitempty" url:"custom"`    //option, 商户的userId
	//sdk设置
	//IPNUrl     string `json:"ipn_url,omitempty" mapstructure:"ipn_url,omitempty" url:"ipn_url"`             // 实际回调地址（ajax）
}

type CoinPayDepositCommonResponse struct {
	Error string `json:"error" mapstructure:"error"`
}

type CoinPayDepositResponse struct {
	Error  string                    `json:"error" mapstructure:"error"`
	Result *CoinPaymentDepositResult `json:"result,omitempty" mapstructure:"result,omitempty"`
}

type CoinPaymentDepositResult struct {
	Amount         string `json:"amount" mapstructure:"amount"`
	Address        string `json:"address" mapstructure:"address"`
	DestTag        string `json:"dest_tag,omitempty" mapstructure:"dest_tag,omitempty"` // 有些币种需要目标标签
	TxnID          string `json:"txn_id" mapstructure:"txn_id"`
	ConfirmsNeeded string `json:"confirms_needed" mapstructure:"confirms_needed"`
	Timeout        int    `json:"timeout" mapstructure:"timeout"`
	CheckoutURL    string `json:"checkout_url" mapstructure:"checkout_url"`
	StatusURL      string `json:"status_url" mapstructure:"status_url"`
	QRCodeURL      string `json:"qrcode_url" mapstructure:"qrcode_url"`
}

type CoinPayCommonBackReq struct {
	IpnType string `json:"ipn_type" mapstructure:"ipn_type"` //消息类型，api是充值,withdrawal是提现
}

// https://www.coinpayments.net/merchant-tools-ipn
type CoinPayDepositBackReq struct {
	IpnType    string `json:"ipn_type" mapstructure:"ipn_type"`                                 //消息类型，写死:api
	Status     string `json:"status" mapstructure:"status"`                                     //支付状态,>=100就是成功!!! 0-pending, 100-confirm/complete
	StatusText string `json:"status_text" mapstructure:"status_text"`                           //支付状态的描述
	TxnID      string `json:"txn_id" mapstructure:"txn_id"`                                     //txId
	Currency1  string `json:"currency1" mapstructure:"currency1"`                               //支付的货币
	Currency2  string `json:"currency2" mapstructure:"currency2"`                               //支付的货币
	Amount1    string `json:"amount1" mapstructure:"amount1"`                                   //总量
	Amount2    string `json:"amount2" mapstructure:"amount2"`                                   // amount in satoshis
	Fee        string `json:"fee" mapstructure:"fee"`                                           //The fee on the payment in the buyer's selected coin.
	Invoice    string `json:"invoice,omitempty" mapstructure:"invoice,omitempty" url:"invoice"` //option, 商户订单号
	Custom     string `json:"custom,omitempty" mapstructure:"custom,omitempty" url:"custom"`    //option, 商户的userId

}

//===========withdraw===================================

type CoinPayWithdrawalRequest struct {
	Amount   string `json:"amount" mapstructure:"amount"`
	Currency string `json:"currency" mapstructure:"currency"` //The cryptocurrency to withdraw. (BTC, LTC, etc.)
	//Currency2   string `json:"currency2" mapstructure:"currency2"` //The cryptocurrency to withdraw. (BTC, LTC, etc.)
	Address string `json:"address" mapstructure:"address"` //提现地址
	//sdk设置
	//IPNUrl     string `json:"ipn_url,omitempty" mapstructure:"ipn_url,omitempty" url:"ipn_url"`             // 实际回调地址（ajax）
	AutoConfirm int `json:"auto_confirm" mapstructure:"auto_confirm"` //设置为1, If set to 1, withdrawal will complete without email confirmation.

	Note string `json:"note,omitempty" mapstructure:"note"` //里边放:商户的订单号
}

type CoinPayWithdrawalCommonResponse struct {
	Error string `json:"error" mapstructure:"error"`
}

// WithdrawalResponse is the response we expect from the API server.
type CoinPayWithdrawalResponse struct {
	Error  string                   `json:"error" mapstructure:"error"`
	Result *CoinPayWithdrawalResult `json:"result" mapstructure:"result"`
}

type CoinPayWithdrawalResult struct {
	Amount string `json:"amount" mapstructure:"amount"`
	ID     string `json:"id" mapstructure:"id"`         //这个是psp的订单号.
	Status int    `json:"status" mapstructure:"status"` // 0 or 1. 0 = transfer created, waiting for email conf. 1 = transfer created with no email conf.
}

// https://www.coinpayments.net/merchant-tools-ipn
type CoinPayWithdrawalBackReq struct {
	IpnType    string `json:"ipn_type" mapstructure:"ipn_type"` //消息类型，写死:withdrawal
	ID         string `json:"id" mapstructure:"id"`             //这个是psp的订单号
	Status     int    `json:"status" mapstructure:"status"`     //状态, 枚举:<0 = failed, 0 = waiting email confirmation, 1 = pending, and 2 = sent/complete.
	StatusText string `json:"status_text" mapstructure:"status_text"`
	Address    string `json:"address" mapstructure:"address"`         //提现地址
	TxnID      string `json:"txn_id,omitempty" mapstructure:"txn_id"` //txID
	Currency   string `json:"currency" mapstructure:"currency"`
	Amount     string `json:"amount" mapstructure:"amount"`
	Amounti    string `json:"amounti" mapstructure:"amounti"`     //The total amount of the withdrawal in Satoshis
	Note       string `json:"note,omitempty" mapstructure:"note"` //里边登记放:商户的订单号
}
