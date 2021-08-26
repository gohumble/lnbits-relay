package main

import "github.com/gohumble/lnbits-relay/lnbits"

func main() {
	relay := lnbits.NewDonationRelay(
		Configuration.Host,
		lnbits.NewClient(Configuration.LnbitsInvoiceKey, Configuration.LnbitsUrl))
	relay.Start()
}
