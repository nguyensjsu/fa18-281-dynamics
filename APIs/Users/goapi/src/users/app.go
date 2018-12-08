package main

import (
	"encoding/json"
	"fmt"
	"log"
	http "net/http"

	. "users/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server1 = "mongodb://admin:admin@54.153.82.51:27017,13.52.93.108:27017,52.9.115.13:27017,50.18.201.231:27017,13.52.91.223:27017"
var mongodb_database 	= "shayona-store"
var mongodb_collection 	= "users"

func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shayona Grocery Store API is alive!")
}

// POST /users: create a new user
func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()

	session, err := mgo.Dial(mongodb_server1)

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)

	err = c.Insert(user)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

// GET /users - get all user
func GetAllUsersEndPoint(w http.ResponseWriter, r *http.Request) {

	session, err := mgo.Dial(mongodb_server1)

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)

	var users []User
	err = c.Find(bson.M{}).All(&users)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

// GET /users/{username}
func GetUserEndPoint(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	session, err := mgo.Dial(mongodb_server1)

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)

	var user User
	err = c.Find(bson.M{"username": params["username"]}).One(&user)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Username")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

// DELETE an existing user
func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	session, err := mgo.Dial(mongodb_server1)

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	err = c.Remove(bson.M{"username": user.Username})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", GetAllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users/{username}", GetUserEndPoint).Methods("GET")
	r.HandleFunc("/users", DeleteUserEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)); err != nil {
		log.Fatal(err)
	}
}
