package go_coinpay

import (
	"github.com/asaka1234/go-coinpay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string
	PrivateKey string
	PublicKey  string
	IPNSecret  string

	EndPoint        string //所有请求都是到这里,通过参数区分
	DepositBackUrl  string //回调地址
	WithdrawBackUrl string //回调地址

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantId string, privateKey string, publicKey, ipnSecret string, endPoint string, depositBackUrl, withdrawBackUrl string) *Client {
	return &Client{
		MerchantID: merchantId,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		IPNSecret:  ipnSecret,

		EndPoint:        endPoint,
		DepositBackUrl:  depositBackUrl,
		WithdrawBackUrl: withdrawBackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
