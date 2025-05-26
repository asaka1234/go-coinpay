package go_buy365

import (
	"github.com/asaka1234/go-buy365/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string
	AccessKey  string
	BackKey    string

	DepositURL         string
	WithdrawURL        string
	WithdrawConfirmURL string
	OrderListURL       string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantId string, accessKey string, backKey string, depositURL string, withdrawURL, withdrawConfirmURL, orderListURL string) *Client {
	return &Client{
		MerchantID: merchantId,
		AccessKey:  accessKey,
		BackKey:    backKey,

		DepositURL:         depositURL,
		WithdrawURL:        withdrawURL,
		WithdrawConfirmURL: withdrawConfirmURL,
		OrderListURL:       orderListURL,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
