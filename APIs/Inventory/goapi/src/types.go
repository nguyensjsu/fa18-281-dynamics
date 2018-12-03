package main

type inventoryItem struct {
	id             	int 	
	name			string    	
	inventory 	 	int	    
	quantity	 	int	
}

type order struct {
	id             	string 	
	OrderStatus 	string	
}

var orders map[string] order

