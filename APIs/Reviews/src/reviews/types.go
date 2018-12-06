package main

import "gopkg.in/mgo.v2/bson"

type Review struct {
	id       	bson.ObjectId
	Review      string
	ItemId		int
	UserId 		int
}