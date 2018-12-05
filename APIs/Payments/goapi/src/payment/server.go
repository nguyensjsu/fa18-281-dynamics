package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	Name		string 	`json:"name" bson:"name"`
	Quantity	int 	`json:"quantity" bson:"quantity"`
	Size		string 	`json:"size" bson:"size"`
	Price		float64 `json:"price" bson:"price"`
}

type Purchase struct {
	Id 			string 	`json:"id" bson:"_id"`
	User 		string 	`json:"user" bson:"user"`
	TotalItems 	int 	`json:"total_items" bson:"total_items"`
	TotalCost 	float64 `json:"total_cost" bson:"total_cost"`
	Cart 		[]Item  `json:"cart" bson:"cart"`
}

// MongoDB Config
//var mongodb_server = "admin:cmpe281@ip-10-0-1-207.us-west-1.compute.internal:27017"
var mongodb_server = "ec2-54-193-109-132.us-west-1.compute.amazonaws.com:27017"
var mongodb_database = "shayona"
var mongodb_collection = "purchases"

// NewServer configures and returns a Server.
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

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payments", getPaymentsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payment", paymentHandler(formatter)).Methods("POST")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Purchase API version 1.0 alive!"})
	}
}

// API Payments Handler - Get all purchases
func getPaymentsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var purchases []bson.M
		err = c.Find(nil).All(&purchases)
		if err != nil {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"No purchases yet!"})
		} else {
			fmt.Println("All purchases: ", purchases)
			formatter.JSON(w, http.StatusOK, purchases)
		}
	}
}

// API Payment Handler - Insert a new payment
func paymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var totalItems int
		var totalCost float64

		decoder := json.NewDecoder(req.Body)
		var t Purchase
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		} else {
			for _, item := range t.Cart {
				totalItems += item.Quantity
				totalCost += float64(item.Quantity) * item.Price
			}
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		uuid, _ := uuid.NewV4()
		entry := Purchase{uuid.String(),
				t.User,
				totalItems,
				math.Floor(totalCost*100)/100,
				t.Cart}
		err = c.Insert(entry)
		if err != nil {
			fmt.Println("Error while inserting purchase: ", err)
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Purchase added"})
		}

	}
}
