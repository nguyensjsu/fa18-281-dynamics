package main

type InventoryItem struct {
	Item_id 			int     `json:"Item_id" bson:"Item_id"`
	Item_name     		string	`json:"Item_name" bson:"Item_name"`
	Item_description 	string	`json:"Item_description" bson:"Item_description"`
	Item_inventory 		int 	`json:"Item_inventory" bson:"Item_inventory"`
	Item_rate    		float64 `json:"Item_rate" bson:"Item_rate"`
}

type Item struct {
	ItemName		string 	`json:"item_name" bson:"item_name"`
	ItemQuantity	int 	`json:"item_quantity" bson:"item_quantity"`
	Rate			float64 `json:"item_rate" bson:"item_rate"`
}

type ItemCart struct {
	Id 			string 	`json:"_id" bson:"_id"`
	Items 		[]Item  `json:"items" bson:"items"`
}

