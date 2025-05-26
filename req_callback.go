package go_buy365

import (
	"errors"
	"github.com/asaka1234/go-buy365/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCancelCallback(req Buy365DepositCancelBackReq, processor func(Buy365DepositCancelBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySignDeposit(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositSucceedCallBack(req Buy365DepositSucceedBackReq, processor func(Buy365DepositSucceedBackReq) error) error {
	//验证签名
	params := map[string]interface{}{
		"bill_no": req.BillNo, //只是value的拼接
	}

	verifyResult := utils.VerifySignWithdraw(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

//==========================================

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCancelCallBack(req Buy365WithdrawCancelBackReq, processor func(Buy365WithdrawCancelBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	verifyResult := utils.VerifySignDeposit(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}

// 充值的回调处理(传入一个处理函数)
func (cli *Client) WithdrawSucceedCallBack(req Buy365WithdrawSucceedBackReq, processor func(Buy365WithdrawSucceedBackReq) error) error {
	//验证签名
	params := map[string]interface{}{
		"bill_no": req.BillNo, //只是value的拼接
	}

	verifyResult := utils.VerifySignWithdraw(params, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}
	if req.SysNo != cli.MerchantID {
		return errors.New("merchanID is wrong!")
	}

	//开始处理
	return processor(req)
}
