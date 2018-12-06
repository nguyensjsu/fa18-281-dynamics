package main

type InventoryItem struct {
	Item_id 			int     `json:"item_id" bson:"item_id"`
	Item_name     		string	`json:"item_name" bson:"item_name"`
	Item_description 	string	`json:"item_description" bson:"item_description"`
	Item_inventory 		int 	`json:"item_inventory" bson:"item_inventory"`
	Item_rate    		float64 `json:"item_rate" bson:"item_rate"`
}

