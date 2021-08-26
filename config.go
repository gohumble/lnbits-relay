package main

import "github.com/jinzhu/configor"

var Configuration = struct {
	LnbitsInvoiceKey string `json:"lnbits_invoice_key" yaml:"lnbits_invoice_key"`
	LnbitsUrl        string `json:"lnbits_url" yaml:"lnbits_url"`
	Host             string `json:"host" yaml:"host"`
}{}

func init() {
	err := configor.Load(&Configuration, "config.yaml")
	if err != nil {
		panic(err)
	}
}
