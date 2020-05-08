package main

//deprecated
func DdcOrderGenerateOrder() {
	Output("DdcOrderGenerateOrder")
	url := "service01:8085/v3/order/generateOrder"
	CurlGET(url)
}

//*	*
func PtgRemoveExpiredInvitation() {
	Output("PtgRemoveExpiredInvitation")
	url := "service01:8088/v3/patungan/removeExpiredInvitation"
	CurlGET(url)
}

//*	*
func PtgRemoveExpiredCheckout() {
	Output("PtgRemoveExpiredCheckout")
	url := "service01:8088/v3/patungan/removeExpiredCheckout"
	CurlGET(url)
}

//*	*
func PtgLockInitiator() {
	Output("PtgLockInitiator")
	url := "service01:8088/v3/patungan/lockInitiator"
	CurlGET(url)
}

//*	*
func CancelCustomDelivery() {
	Output("CancelCustomDelivery")
	url := "service01:8088/v3/order/cancelcustomdelivery/1/120"
	CurlGET(url)
}

//*/5       *
func VoucherRemoveReserveAndHistory() {
	Output("VoucherRemoveReserveAndHistory")
	url := "service01:8088/v3/voucher/removeReserveAndHistory"
	CurlGET(url)
}

//*       *
func VoucherCreateHistory() {
	Output("VoucherCreateHistory")
	url := "service01:8088/v3/voucher/createHistory"
	CurlGET(url)
}

//*       *
func VoucherCloseExpired() {
	Output("VoucherCloseExpired")
	url := "service01:8088/v3/voucher/closeExpired"
	CurlGET(url)
}

//55 23
func ProductNumberOfSales() {
	Output("ProductNumberOfSales")
	url := "service01:8088/v3/product/numberofsales"
	CurlGET(url)
}

//59 23
func ProductReview() {
	Output("ProductReview")
	url := "service01:8088/v3/product/review"
	CurlGET(url)
}

//*/5     *
func PromoCloseExpired() {
	Output("PromoCloseExpired")
	url := "service01:8088/v3/promo/closeExpired"
	CurlGET(url)
}

//*/10    *
func CancelOrder11() {
	Output("Cancel Order 1 1")
	url := "service01:8085/v3/order/cancel_orders/1/1"
	CurlGET(url)
}

//*/10    *
func CancelOrder12() {
	Output("Cancel Order 1 2")
	url := "service01:8085/v3/order/cancel_orders/1/2"
	CurlGET(url)
}

//*/10    *
func CancelOrder16() {
	Output("Cancel Order 1 6")
	url := "service01:8085/v3/order/cancel_orders/1/6"
	CurlGET(url)
}

//*/10    *
func CancelOrder14() {
	Output("Cancel Order 1 4")
	url := "service01:8085/v3/order/cancel_orders/1/4"
	CurlGET(url)
}

//59 	23
func CheckMemberExpired() {
	Output("CheckMemberExpired")
	url := "service01:8088/v3/member/checkExpireDate"
	CurlGET(url)
}

//59 	23
func CheckMemberExpiredPoint() {
	Output("CheckMemberExpiredPoint")
	url := "service01:8088/v3/member/checkExpirePointDate"
	CurlGET(url)
}

//00 	15
func EmailStepExpireDate() {
	Output("EmailStepExpireDate")
	url := "service01:8088/v3/member/emailStepExpireDate/-3"
	CurlGET(url)
}

//5	0
func BoCheckPublish() {
	Output("BoCheckPublish")
	url := "backoffice.ddc.vm/api/checkPublish"
	CurlGET(url)
}

//00      15
func EventClosed() {
	Output("EventClosed")
	url := "service01:8088/v3/event/eventclosed"
	CurlGET(url)
}

//0     *
func ProductCalcMosAll() {
	Output("ProductCalcMosAll")
	url := "service01:8088/v3/product/calculatemosall"
	CurlGET(url)
}

//0     *
func ProductCalcMosHdm() {
	Output("ProductCalcMosHdm")
	url := "service01:8088/v3/product/calculatemoshdm"
	CurlGET(url)
}

//0     1
func ProductStockOpname() {
	Output("ProductStockOpname")
	url := "service01:8088/v3/product/stockopname"
	CurlGET(url)
}
