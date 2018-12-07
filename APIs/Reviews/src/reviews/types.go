package main

import "gopkg.in/mgo.v2/bson"

type Review struct {
	Id      	bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ItemName	string

	Reviews []struct {
		ReviewName 		string
		ReviewSummary 	string
		Review 			string
		Rating 			int
	}
}