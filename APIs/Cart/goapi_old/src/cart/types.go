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
	User_Name string `json:"User_Name"`

	Items []struct {
		Name   string  `json:"name"`
		Count  int     `json:"count"`
		Price   float64 `json:"Price"`
		Amount float64 `json:"amount"`
	} `json:"items"`

	Total float64 `json:"total"`
}

type Keys struct {
	Keys []string
}

type Order struct {
	OrderId string
}
