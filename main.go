package main

import "github.com/gohumble/lnbits-relay/lnbits"

func main() {
	relay := lnbits.NewDonationRelay(
		Configuration.Host,
		lnbits.NewClient(Configuration.LnbitsKey, Configuration.LnbitsUrl),
		lnbits.RelayConfiguration{
			AdminKey: Configuration.AdminKey,
			Donee:    Configuration.Donee})
	relay.Start()
}
