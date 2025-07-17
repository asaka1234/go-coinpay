文档
==============
https://www.coinpayments.net/apidoc-create-transaction
https://www.coinpayments.net/merchant-tools-ipn
用的是 legacy 版本

流程
===============
支持的symbol: https://www.coinpayments.net/supported-coins-stable




comment
===============
1. 请求: application/x-www-form-urlencoded


v2
===============
1. deposit:   
   1. 发起:  https://a-docs.coinpayments.net/api/invoices/examples/invoice-link
   2. 回调:  https://a-docs.coinpayments.net/api/webhooks/clients/payload
2. withdrawl
   1. 发起:  https://a-docs.coinpayments.net/api/transactions/routes/postMerchantWalletsSpendRequestV2ById
   2. 回调： https://a-docs.coinpayments.net/api/webhooks/wallets/payload