package main

import "gopkg.in/mgo.v2/bson"

type Review struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Review string
	ItemId int
	UserId int
}