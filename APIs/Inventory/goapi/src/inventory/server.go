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
var mongodb_server = "mongodb://admin:admin@52.9.134.184:27017/admin"
var mongodb_database = "inventory"
var mongodb_collection = "inventoryItem"

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
	mx.HandleFunc("/addInventoryItem", addItemHandler(formatter)).Methods("POST")
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



// API Post (Add Item to DB) Handler.
func addItemHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var m inventoryItem
		_ = json.NewDecoder(req.Body).Decode(&m)
		fmt.Println("Item is: ", m, " ", m.name," ",m.desc," ",m.quantity," ", m.price)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		entry := inventoryItem{
			name: m.name,
			desc: m.desc,
			quantity: m.quantity,
			price: m.price,
		}
		err = c.Insert(entry)
		if err != nil {
			fmt.Println("Error while adding item to inventory: ", err)
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Item added to inventory"})
		}
	}
}


//sample document
// db.createCollection("inventoryItem")

// db.inventoryItem.insert({'id':1,'name':'rice','desc':'This is rice','quantity':5,'price':15})

// use admin
// db.createUser( {
// 	user: "admin",
// 	pwd: "admin",
// 	roles: [{ role: "root", db: "admin" }]
// });

//const IP_MONGODB_DATABASE = 'mongodb://admin:admin@52.53.77.103:27017/admin';
