package main

type Wallet struct {
	Username 	string 	`json:"username" bson:"username"`
	Amount		float64 `json:"wallet_amount" bson:"wallet_amount"`
}

type Item struct {
	ItemName		string 	`json:"item_name" bson:"item_name"`
	ItemQuantity	int 	`json:"item_quantity" bson:"item_quantity"`
	Rate			float64 `json:"item_rate" bson:"item_rate"`
}

type Purchase struct {
	Id 			string 	`json:"_id" bson:"_id"`
	Username 	string 	`json:"username" bson:"username"`
	TotalItems 	int 	`json:"item_count" bson:"item_count"`
	CartTotal 	float64 `json:"cart_total" bson:"cart_total"`
	Items 		[]Item  `json:"items" bson:"items"`
}
