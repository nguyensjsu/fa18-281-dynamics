package main


import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)


// MongoDB Config
var mongodb_server = "mongodb://admin:admin@10.0.5.148:27017,10.0.6.227:27017,10.0.5.134:27017,10.0.6.219:27017,10.0.5.13:27017"
// var mongodb_server = "mongodb://admin:admin@13.52.14.112"
var mongodb_database = "inventory"
var mongodb_collection = "InventoryItem"

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/inventory", getInventoryHandler(formatter)).Methods("GET")
	mx.HandleFunc("/inventory", addItemToInventoryHandler(formatter)).Methods("POST")
	mx.HandleFunc("/inventory/update", updateInventoryHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/inventory/delete/item_name", deleteInventoryByItem(formatter)).Methods("DELETE")
	mx.HandleFunc("/inventory/delete", deleteInventoryHandler(formatter)).Methods("DELETE")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Info string }{"Ping works!"})
	}
}

// API GetInventoryItems Handler -- get all the items from the inventory
func getInventoryHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var results []bson.M
		err = c.Find(bson.M{}).All(&results)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
		formatter.JSON(w, http.StatusOK, results)
	}
}




// API update item inventory - decrement the item inventory after processing orders
func updateInventoryHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var i ItemCart
		err := decoder.Decode(&i)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		} else {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var actual_inventory bson.M
			for _, item := range i.Items {
				fmt.Println("Display Inv ", item.ItemName)
				sub_quantity := item.ItemQuantity
				err = c.Find(bson.M{"Item_name" : item.ItemName}).One(&actual_inventory)		
				actual_quantity := actual_inventory["Item_inventory"].(int) - int(sub_quantity)
				query := bson.M{"Item_name" : item.ItemName}
		        change := bson.M{"$set": bson.M{ "Item_inventory" : actual_quantity}}
		        err = c.Update(query, change)
		        if err != nil {
		                log.Fatal(err)
		                formatter.JSON(w, http.StatusOK, struct{ Info string }{"Update failed"})
		        } 
				}
				formatter.JSON(w, http.StatusOK, struct{ Info string }{"Updated inventory items successfully"})
			}

		}
	}


// API Add an item to inventory - Add an item to the inventory
func addItemToInventoryHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var i InventoryItem
		_ = json.NewDecoder(req.Body).Decode(&i)
		fmt.Println("Item is: ", i.Item_name)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		entry := InventoryItem{
			Item_id: i.Item_id,
			Item_name: i.Item_name,
			Item_description: i.Item_description,
			Item_inventory: i.Item_inventory,
			Item_rate: i.Item_rate,
		}
		err = c.Insert(entry)
		if err != nil {
			fmt.Println("Error while adding item: ", err)
			formatter.JSON(w, http.StatusInternalServerError, struct{ Response error }{err})
		} else {
			formatter.JSON(w, http.StatusNoContent, struct{ Response string }{"Item added"})
		}
	}
}


// API delete All Item - Delete all Inventory Items
func deleteInventoryHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var i InventoryItem
		err := decoder.Decode(&i)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		_, err = c.RemoveAll(bson.M{})
		if err != nil {
			formatter.JSON(w, http.StatusOK, struct{ Info string }{"No items in the inventory"})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Info string }{"All the items in the inventory deleted"})
		}
	}
}

// API Delete Inventory for a specified item
func deleteInventoryByItem(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var i InventoryItem
		err := decoder.Decode(&i)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		err = c.Remove(bson.M{"Item_name":i.Item_name})
		if err != nil {
			formatter.JSON(w, http.StatusOK, struct{ Info string }{"No items in the inventory with this item name"})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Info string }{"Item deleted"})
		}
	}
}





//sample document
// db.createCollection("InventoryItem")

// db.InventoryItem.insert({'Item_id':1,'Item_name':'milk','Item_description':'This is milk','Item_inventory':20,'Item_rate':15})
// db.InventoryItem.insert({'Item_id':1,'Item_name':'yogurt','Item_description':'This is yogurt','Item_inventory':100,'Item_rate':12})

// use admin
// db.createUser( {
// 	user: "admin",
// 	pwd: "admin",
// 	roles: [{ role: "root", db: "admin" }]
// });

//const IP_MONGODB_DATABASE = 'mongodb://admin:admin@52.53.77.103:27017/admin';
// mongo -u admin -p admin --authenticationDatabase admin





