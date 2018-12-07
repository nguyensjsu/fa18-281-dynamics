package main

type InventoryItem struct {
	Item_id 			int     `json:"item_id" bson:"item_id"`
	Item_name     		string	`json:"item_name" bson:"item_name"`
	Item_description 	string	`json:"item_description" bson:"item_description"`
	Item_inventory 		int 	`json:"item_inventory" bson:"item_inventory"`
	Item_rate    		float64 `json:"item_rate" bson:"item_rate"`
	Item_quantity  		int 	`json:"item_quantity" bson:"item_quantity"`
	Item_subtotal		float64	`json:"item_subtotal" bson:"item_subtotal"`

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

