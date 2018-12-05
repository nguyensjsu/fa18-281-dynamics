package dao

import (
	"log"

	. "users/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
	Username string
	Password string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

// Establish a connection to a database
func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal("MongoDB Error: ", err)
	}
	db = session.DB("shayona-store")
}

// Insert a user into database
func (m *UsersDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Find all users
func (m *UsersDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Find a user by its username
func (m *UsersDAO) FindByUsername(username string) (User, error) {
	var user User
	err := db.C(COLLECTION).Find(bson.M{"username": username}).One(&user)

	return user, err
}
