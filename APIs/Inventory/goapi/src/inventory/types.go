package main

import "gopkg.in/mgo.v2/bson"

type inventoryItem struct {
	id       	bson.ObjectId
	name     	string	
	desc     	string	
	quantity 	int 	
	price    	int
}
