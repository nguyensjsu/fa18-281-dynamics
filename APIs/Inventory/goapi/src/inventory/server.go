package main


import (
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

		if err != nil {
			fmt.Println("Inventory API - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var results []inventoryItem
		err = c.Find(bson.M{}).All(&results)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
		formatter.JSON(w, http.StatusOK, results)
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
