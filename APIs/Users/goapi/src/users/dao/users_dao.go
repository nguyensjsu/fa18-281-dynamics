package dao

import (
	"log"

	. "users/models"

	mgo "gopkg.in/mgo.v2"
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
