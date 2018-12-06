package main

type Review struct {
	Id     string `json:"_id" bson:"_id"`
	Review string
	ItemId int
	UserId int
}