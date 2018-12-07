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


// MongoDB Config
var mongodb_server = "mongodb://admin:cmpe281@52.53.82.217:27017,54.177.200.126:27017,13.52.64.28:27017/groupproject?authSource=admin&replicaSet=cmpe281"
var mongodb_database = "groupproject"
var mongodb_collection = "reviews"


// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/getReviews", getReviewsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/postReview", postReviewHandler(formatter)).Methods("POST")
	mx.HandleFunc("/updateReview", updateReviewHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/deleteReview", deleteReviewHandler(formatter)).Methods("DELETE")

}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Ping WORKS!!!"})
	}
}

// API Get All Reviews Handler
func getReviewsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Reviews API (Get) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var results []Review
		err = c.Find(bson.M{}).All(&results)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(results)
		formatter.JSON(w, http.StatusOK, results)
	}
}

// API Post a Review Handler.
func postReviewHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var m Review
		_ = json.NewDecoder(req.Body).Decode(&m)
		fmt.Println("Review is: ", m.ItemName, " ", m.Reviews)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Reviews API (Post) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		entry := Review{
			ItemName: m.ItemName,
			Reviews: m.Reviews,
		}
		err = c.Insert(entry)
		if err != nil {
			fmt.Println("Error while adding reviews: ", err)
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Review added"})
		}
	}
}


// API Update a Review Handler.
func updateReviewHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var m Review
		_ = json.NewDecoder(req.Body).Decode(&m)
		fmt.Println("Review is: ", m.ItemName, " " , "Reviews", m.Reviews)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Reviews API (Update) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		query := bson.M{
			"itemname": m.ItemName,
		}
		change := bson.M{"$push": bson.M{ "reviews" : bson.M{"$each": m.Reviews }}}
		err = c.Update(query, change)

		if err != nil {
			fmt.Println("Error while updating reviews: ", err)
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Review updated"})
		}
	}
}


// API Delete a Review Handler.
func deleteReviewHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var m Review
		_ = json.NewDecoder(req.Body).Decode(&m)
		fmt.Println("Review is: ", m.Reviews)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		if err != nil {
			fmt.Println("Reviews API (Delete) - Unable to connect to MongoDB during read operation")
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		query := bson.M{
			"itemname": m.ItemName,
		}
		change := bson.M{"$pull": bson.M{ "reviews" :  bson.M{"$in": m.Reviews } }}
		err = c.Update(query, change)

		if err != nil {
			fmt.Println("Error while deleting reviews: ", err)
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Review Deleted"})
		}
	}
}
