package main

func DdcOrderGenerateOrder() {
	Output("DdcOrderGenerateOrder")
	url := "service01:8085/v3/order/generateOrder"
	CurlGET(url)
}

func RemoveExpiredInvitation() {
	Output("RemoveExpiredInvitation")
	url := "service01:8088/v3/patungan/removeExpiredInvitation"
	CurlGET(url)
}

func RemoveExpiredCheckout() {
	Output("RemoveExpiredCheckout")
	url := "service01:8088/v3/patungan/removeExpiredCheckout"
	CurlGET(url)
}
