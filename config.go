package main

import "github.com/jinzhu/configor"

var Configuration = struct {
	AdminKey  string `json:"lnbits_admin_id" yaml:"lnbits_admin_id"`
	LnbitsKey string `json:"lnbits_admin_key" yaml:"lnbits_admin_key"`
	LnbitsUrl string `json:"lnbits_url" yaml:"lnbits_url"`
	Host      string `json:"host" yaml:"host"`
	Donee     string `json:"donee_id" yaml:"donee_id"`
}{}

func init() {
	err := configor.Load(&Configuration, "config.yaml")
	if err != nil {
		panic(err)
	}
}
