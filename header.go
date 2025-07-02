package go_coinpay

func getHeaders(hmac string) map[string]string {
	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"HTTP_HMAC":         hmac,
	}
}
