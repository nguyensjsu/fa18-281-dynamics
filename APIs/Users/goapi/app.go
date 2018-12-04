package main

import (
	"encoding/json"
	"fmt"
	"log"
	http "net/http"

	. "./config"
	. "./dao"
	. "./models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = UsersDAO{}

func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shayona Grocery Store API is alive!")
}

// POST: create a new user
func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := dao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}
