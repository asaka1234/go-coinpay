package go_coinpay

import (
	"github.com/asaka1234/go-coinpay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params CoinPayInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params CoinPayInitParams) *Client {
	return &Client{
		Params: params,
		
		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
