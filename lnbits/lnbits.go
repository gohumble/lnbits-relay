package lnbits

import (
	"github.com/imroc/req"
)

// NewClient returns a new lnbits api client. Pass your API key and url here.
func NewClient(key, url string) *Client {
	return &Client{
		url: url,
		header: req.Header{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"X-Api-Key":    key,
		},
	}
}

// GetUser returns user information
func (c *Client) GetUser(userId string) (user User, err error) {
	resp, err := req.Post(c.url+"/usermanager/api/v1/users/"+userId, c.header, nil)
	if err != nil {
		return
	}

	if resp.Response().StatusCode >= 300 {
		var reqErr Error
		resp.ToJSON(&reqErr)
		err = reqErr
		return
	}

	err = resp.ToJSON(&user)
	return
}

// Invoice creates an invoice associated with this wallet.
func (c Client) Invoice(params InvoiceParams, w Wallet) (lntx Invoice, err error) {
	resp, err := req.Post(c.url+"/api/v1/payments", w.header, req.BodyJSON(&params))
	if err != nil {
		return
	}
	if resp.Response().StatusCode >= 300 {
		var reqErr Error
		resp.ToJSON(&reqErr)
		err = reqErr
		return
	}

	err = resp.ToJSON(&lntx)
	return
}

// Wallets returns all wallets belonging to an user
func (c Client) Wallets(w User) (wtx []Wallet, err error) {
	resp, err := req.Get(c.url+"/usermanager/api/v1/wallets/"+w.ID, c.header, nil)
	if err != nil {
		return
	}

	if resp.Response().StatusCode >= 300 {
		var reqErr Error
		resp.ToJSON(&reqErr)
		err = reqErr
		return
	}

	err = resp.ToJSON(&wtx)
	return
}
