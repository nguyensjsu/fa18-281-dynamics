package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// MongoDB Config
var mongodb_server = "mongodb://admin:admin@52.9.134.184:27017/inventory?authSource=admin"
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

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Inventory API is live!!"})
		mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	}
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
}

// Ping the API to check if its working.
func (c *Client) Ping() (string, error) {
	resp, err := c.Get(c.Endpoint + "/ping")

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return "Ping Error!", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if debug {
		fmt.Println("[RIAK DEBUG] GET: " + c.Endpoint + "/ping => " + string(body))
	}
	return string(body), nil
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
