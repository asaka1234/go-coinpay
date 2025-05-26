package go_buy365

const (
	MERCHANT_ID = "502326"                               //商户号
	ACCESS_KEY  = "2299de9c55458f1d48611cdd9073aa7c"     //调用psp的签名key
	BACK_KEY    = "5E1572C3-0274-A6C7-E797-6E3111CCDC71" //回调的签名key

	//deposit
	DEPOSIT_URL = "https://swpapi.fastgo788.io/UtInRecordApi/orderGateWay"

	//withdraw
	WITHDRAW_URL         = "https://mmapi.qhcm12.com/AjaxOpen/saveOutOrder"
	WITHDRAW_CONFIRM_URL = "https://mmapi.proxima131.com/AjaxOpen/appealOutOrder"

	//orderlist
	ORDERLIST_URL = "https://mmapi.proxima131.com/AjaxOpen/getOutOrderList"
)
