package lnbits

import (
	"github.com/imroc/req"
)

type Client struct {
	header     req.Header
	url        string
	AdminKey   string
	InvoiceKey string
}

type User struct {
	ID     string `json:"id"`
	Wallet *Wallet
}

type InvoiceParams struct {
	Out     bool   `json:"out"`
	Amount  int64  `json:"amount"`
	Memo    string `json:"memo"` // the invoice description.
	Webhook string `json:"webhook,omitempty"`
}

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"status"`
}

func (err Error) Error() string {
	return err.Message
}

type Wallet struct {
	*Client  `gorm:"-"`
	ID       string `json:"id" gorm:"id"`
	Adminkey string `json:"adminkey"`
	Inkey    string `json:"inkey"`
	Balance  int64  `json:"balance"`
	Name     string `json:"name"`
	User     string `json:"user"`
}
type Invoice struct {
	PaymentHash    string `json:"payment_hash"`
	PaymentRequest string `json:"payment_request"`
}
