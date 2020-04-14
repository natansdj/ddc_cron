package main

func DdcOrderGenerateOrder() {
	Output("DdcOrderGenerateOrder")
	url := "service01:8085/v3/order/generateOrder"
	CurlGET(url)
}

func PtgRemoveExpiredInvitation() {
	Output("PtgRemoveExpiredInvitation")
	url := "service01:8088/v3/patungan/removeExpiredInvitation"
	CurlGET(url)
}

func PtgRemoveExpiredCheckout() {
	Output("PtgRemoveExpiredCheckout")
	url := "service01:8088/v3/patungan/removeExpiredCheckout"
	CurlGET(url)
}

func CancelCustomDelivery() {
	Output("CancelCustomDelivery")
	url := "service01:8088/v3/order/cancelcustomdelivery/1/120"
	CurlGET(url)
}

func VoucherRemoveReserveAndHistory() {
	Output("VoucherRemoveReserveAndHistory")
	url := "service01:8088/v3/voucher/removeReserveAndHistory"
	CurlGET(url)
}

func VoucherCreateHistory() {
	Output("VoucherCreateHistory")
	url := "service01:8088/v3/voucher/createHistory"
	CurlGET(url)
}

func VoucherCloseExpired() {
	Output("VoucherCloseExpired")
	url := "service01:8088/v3/voucher/closeExpired"
	CurlGET(url)
}

func PromoCloseExpired() {
	Output("PromoCloseExpired")
	url := "service01:8088/v3/promo/closeExpired"
	CurlGET(url)
}

func CancelOrder11() {
	Output("Cancel Order 1 1")
	url := "service01:8085/v3/order/cancel_orders/1/1"
	CurlGET(url)
}

func CancelOrder12() {
	Output("Cancel Order 1 2")
	url := "service01:8085/v3/order/cancel_orders/1/2"
	CurlGET(url)
}

func CancelOrder16() {
	Output("Cancel Order 1 6")
	url := "service01:8085/v3/order/cancel_orders/1/6"
	CurlGET(url)
}

func CancelOrder14() {
	Output("Cancel Order 1 4")
	url := "service01:8085/v3/order/cancel_orders/1/4"
	CurlGET(url)
}

func CheckMemberExpired() {
	Output("CheckMemberExpired")
	url := "service01:8088/v3/member/checkExpireDate"
	CurlGET(url)
}

func CheckMemberExpiredPoint() {
	Output("CheckMemberExpiredPoint")
	url := "service01:8088/v3/member/checkExpirePointDate"
	CurlGET(url)
}
