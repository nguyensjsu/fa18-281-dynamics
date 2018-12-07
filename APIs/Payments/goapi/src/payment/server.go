package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/rs/cors"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Wallet struct {
	Username 	string 	`json:"username" bson:"username"`
	Amount		float64 `json:"wallet_amount" bson:"wallet_amount"`
}

type Item struct {
	ItemName		string 	`json:"item_name" bson:"item_name"`
	ItemQuantity	int 	`json:"item_quantity" bson:"item_quantity"`
	Rate			float64 `json:"item_rate" bson:"item_rate"`
}

type Purchase struct {
	Id 			string 	`json:"_id" bson:"_id"`
	Username 	string 	`json:"username" bson:"username"`
	TotalItems 	int 	`json:"item_count" bson:"item_count"`
	CartTotal 	float64 `json:"cart_total" bson:"cart_total"`
	Items 		[]Item  `json:"items" bson:"items"`
	PaymentInfo string 	`json:"payment_info" bson:"payment_info"`
}

// MongoDB Config
var mongodb_server = "admin:cmpe281@10.0.1.207:27017,10.0.1.217:27017,10.0.1.127:27017,10.0.1.157:27017,10.0.1.160:27017"
var mongodb_database 			= "shayona"
var mongodb_purchase_collection = "purchases"
var mongodb_wallet_collection 	= "wallets"

// NewServer configures and returns a server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
		corsObj := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
    })
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.Use(corsObj)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payments", getPaymentsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payment", paymentHandler(formatter)).Methods("POST")
	mx.HandleFunc("/payments/user", getPaymentsByUserHandler(formatter)).Methods("GET")
	mx.HandleFunc("/payment/delete/id", deletePaymentByIdHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/payments/delete/user", deletePaymentsByUserHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/wallet", getWalletHandler(formatter)).Methods("GET")
	mx.HandleFunc("/wallet", addWalletHandler(formatter)).Methods("POST")
	mx.HandleFunc("/wallet/add", addMoneyToWalletHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/wallet/pay", payWalletHandler(formatter)).Methods("PUT")
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
		c := session.DB(mongodb_database).C(mongodb_purchase_collection)

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

// API Payment Handler - Insert a new purchase after payment
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
			for _, item := range t.Items {
				totalItems += item.ItemQuantity
				totalCost += float64(item.ItemQuantity) * item.Rate
			}
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_purchase_collection)

		uuid, _ := uuid.NewV4()
		entry := Purchase{uuid.String(),
				t.Username,
				totalItems,
				math.Floor(totalCost*100)/100,
				t.Items,
				t.PaymentInfo}
		err = c.Insert(entry)

		if err != nil {
			fmt.Println("Error while inserting purchase: ", err)
		} else {
			jData, _ := json.Marshal(entry)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jData)
		}
	}
}

// API Payments By User Handler - Get all purchases from a specified user
func getPaymentsByUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var t Purchase
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_purchase_collection)

		var purchases []bson.M
		err = c.Find(bson.M{"user":t.Username}).All(&purchases)
		if err != nil {
			formatter.JSON(w, http.StatusNoContent, struct{ Test string }{"No purchases from this user"})
		} else {
			fmt.Println("All purchases: ", purchases)
			formatter.JSON(w, http.StatusOK, purchases)
		}
	}
}


// API Delete Payment By Id Handler - Delete a single payment with specified id
func deletePaymentByIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var t Purchase
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_purchase_collection)

		err = c.Remove(bson.M{"_id":t.Id})

		if err != nil {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"No purchase with this id"})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Purchase deleted"})
		}
	}
}

// API Delete Payments By User Handler - Delete all payments made by a specified user
func deletePaymentsByUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var t Purchase
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_purchase_collection)

		_, err = c.RemoveAll(bson.M{"username":t.Username})
		if err != nil {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"No purchases from this user"})
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Test string }{"Purchases deleted"})
		}
	}
}


// API Get Wallet Handler - Get Wallet for a specified user
func getWalletHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var body Wallet
		err := decoder.Decode(&body)
		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_wallet_collection)

		var wallet []bson.M
		err = c.Find(bson.M{"username":body.Username}).All(&wallet)

		if err != nil {
			fmt.Println("Error searching DB for wallet: ", err)
		} else {
			if (wallet == nil) {
				formatter.JSON(w, http.StatusNoContent, struct{ Test string }{"No wallet for this user"})
			} else {
				fmt.Println("Wallet: ", wallet)
				formatter.JSON(w, http.StatusOK, wallet)
			}

		}
	}
}

func addWalletHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var body Wallet
		err := decoder.Decode(&body)

		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_wallet_collection)

		entry := Wallet{ body.Username, body.Amount }
		err = c.Insert(entry)

		if err != nil {
			fmt.Println("Error while inserting wallet: ", err)
		} else {
			jData, _ := json.Marshal(entry)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jData)
		}
	}
}

func addMoneyToWalletHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var body Wallet
		err := decoder.Decode(&body)

		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_wallet_collection)

		var currentWallet bson.M
		err = c.Find(bson.M{"username" : body.Username}).One(&currentWallet)

		if (err != nil) {
			formatter.JSON(w, http.StatusNoContent, struct{ Test string }{"No wallet for this user"})
		} else {
			currentAmount := currentWallet["wallet_amount"].(float64)
	        query := bson.M{"username" : body.Username}
	        newAmount := currentAmount + body.Amount
	        newAmount = math.Floor(newAmount*100)/100
	        change := bson.M{"$set": bson.M{ "wallet_amount" : newAmount}}
	        err = c.Update(query, change)

			if err != nil {
				log.Fatal(err)
	        } else {
				fmt.Print("Wallet now has $",newAmount,"\n")
				formatter.JSON(w, http.StatusOK, struct{ Test string }{"Wallet amount updated"})
	        }
		}
	}
}

func payWalletHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var body Wallet
		err := decoder.Decode(&body)

		if err != nil {
			fmt.Println("Error parsing the request's body: ", err)
		}

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_wallet_collection)

		var currentWallet bson.M
		err = c.Find(bson.M{"username" : body.Username}).One(&currentWallet)

		if (err != nil) {
			formatter.JSON(w, http.StatusNoContent, struct{ Test string }{"No wallet for this user"})
		} else {
			currentAmount := currentWallet["wallet_amount"].(float64)
	        query := bson.M{"username" : body.Username}
	        newAmount := currentAmount - body.Amount
	        newAmount = math.Floor(newAmount*100)/100
	        change := bson.M{"$set": bson.M{ "wallet_amount" : newAmount}}
	        err = c.Update(query, change)

			if err != nil {
				log.Fatal(err)
	        } else {
				fmt.Print("Wallet now has $",newAmount,"\n")
				formatter.JSON(w, http.StatusOK, struct{ Test string }{"Wallet amount updated"})
	        }
		}
	}
}