package go_coinpay

// ----------pre order-------------------------

// https://www.coinpayments.net/apidoc-create-transaction
type CoinPayDepositReq struct {
	// Required fields
	Cmd        string `json:"cmd" mapstructure:"cmd"`                 // create_transaction - Yes
	Amount     string `json:"amount" mapstructure:"amount"`           // The amount of the payment in the original currency - Yes
	Currency1  string `json:"currency1" mapstructure:"currency1"`     // The original currency of the payment - Yes
	Currency2  string `json:"currency2" mapstructure:"currency2"`     // The currency the buyer will be sending - Yes
	BuyerEmail string `json:"buyer_email" mapstructure:"buyer_email"` // Buyer's email address - Yes

	// Optional fields
	Address    string `json:"address,omitempty" mapstructure:"address,omitempty"`         // Address to send the funds to - No
	BuyerName  string `json:"buyer_name,omitempty" mapstructure:"buyer_name,omitempty"`   // Buyer's name for reference - No
	ItemName   string `json:"item_name,omitempty" mapstructure:"item_name,omitempty"`     // Item name for reference - No
	ItemNumber string `json:"item_number,omitempty" mapstructure:"item_number,omitempty"` // Item number for reference - No
	Invoice    string `json:"invoice,omitempty" mapstructure:"invoice,omitempty"`         // Custom field for your use - No
	Custom     string `json:"custom,omitempty" mapstructure:"custom,omitempty"`           // Custom field for your use - No
	IPNUrl     string `json:"ipn_url,omitempty" mapstructure:"ipn_url,omitempty"`         // URL for IPN callbacks - No
	SuccessUrl string `json:"success_url,omitempty" mapstructure:"success_url,omitempty"` // URL for successful payment - No
	CancelUrl  string `json:"cancel_url,omitempty" mapstructure:"cancel_url,omitempty"`   // URL for cancelled payment - No
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

//===========callback===================================

// https://www.coinpayments.net/apidoc-get-callback-address

type CoinPayDepositBackReq struct {
	// Required fields
	Cmd      string `json:"cmd" mapstructure:"cmd"`           // get_callback_address - Yes
	Currency string `json:"currency" mapstructure:"currency"` // The currency the buyer will be sending - Yes

	// Optional fields
	IPNUrl string `json:"ipn_url,omitempty" mapstructure:"ipn_url,omitempty"` // URL for IPN callbacks - No
	Label  string `json:"label,omitempty" mapstructure:"label,omitempty"`     // Address label - No
	EIP55  int    `json:"eip55,omitempty" mapstructure:"eip55,omitempty"`     // EIP-55 format flag (0/1) - No
}
