/*
	Gumball API in Go
	Uses MySQL & Riak KV
*/

package main

import (
	"net/http"
)

type Client struct {
	Endpoint string
	*http.Client
}

type Cart struct {
	Id     string `json:"id"`
	UserName string `json:"username"`

	Items []struct {
		ItemName   string  `json:"item_name"`
		ItemQuantity  int     `json:"item_quantity"`
		ItemRate   float64 `json:"item_rate"`
		ItemSubtotal float64 `json:"item_subtotal"`
	} `json:"items"`

	CartTotal float64 `json:"cart_total"`
}

type Keys struct {
	Keys []string
}

type Order struct {
	OrderId string
}
