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
var mongodb_server = "mongodb://admin:admin@54.153.77.225"
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
	// mx.HandleFunc("/inventory", addItemToInventoryHandler(formatter)).Methods("POST")
	// mx.HandleFunc("/inventory", updateInventoryHandler(formatter)).Methods("PUT")
	// mx.HandleFunc("/addInventoryItem", addItemHandler(formatter)).Methods("POST")
	mx.HandleFunc("/inventory/delete/item_name", deleteInventoryByItem(formatter)).Methods("DELETE")
	mx.HandleFunc("/inventory/delete", deleteInventoryHandler(formatter)).Methods("DELETE")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Ping works!"})
	}
}

// API GetInventoryItems Handler
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




// // API item to Inventory
// func addItemToInventoryHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {

// 		decoder := json.NewDecoder(req.Body)
// 		var i InventoryItem
// 		err := decoder.Decode(&i)
// 		if err != nil {
// 			fmt.Println("Error parsing the request's body: ", err)
// 		} 
// 		session, err := mgo.Dial(mongodb_server)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer session.Close()
// 		session.SetMode(mgo.Monotonic, true)
// 		c := session.DB(mongodb_database).C(mongodb_collection)

// 		uuid, _ := uuid.NewV4()
// 		entry := InventoryItem{uuid.String(),
// 				i.Item_name,
// 				i.Item_description,
// 				i.Item_id,
// 				i.Item_rate,
// 				i.Item_inventory,
// 			}
// 		err = c.Insert(entry)
// 		if err != nil {
// 			fmt.Println("Error while inserting purchase: ", err)
// 		} else {
// 			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Item added to Inventory"})
// 		}

// 	}
// }


// API Delete all Inventory Items
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
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"No items in the inventory"})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"All the items in the inventory deleted"})
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
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"No items in the inventory with this item name"})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Item deleted"})
		}
	}
}

//sample document
// db.createCollection("InventoryItem")

// db.InventoryItem.insert({'Item_id':1,'Item_name':'milk','Item_description':'This is milk','Item_inventory':5,'Item_rate':15})
// db.InventoryItem.insert({'Item_id':1,'Item_name':'yogurt','Item_description':'This is yogurt','Item_inventory':3,'Item_rate':12})

// use admin
// db.createUser( {
// 	user: "admin",
// 	pwd: "admin",
// 	roles: [{ role: "root", db: "admin" }]
// });

//const IP_MONGODB_DATABASE = 'mongodb://admin:admin@52.53.77.103:27017/admin';
// mongo -u admin -p admin --authenticationDatabase admin


