package database

import (
	"log"

	"github.com/laidingqing/amadd9/common/config"
	mgo "gopkg.in/mgo.v2"
)

var (
	mgoSession   *mgo.Session
	databaseName = config.Database.MainDb
)

func InitDb() {

}

func SetupDb() {

}

//InitDb , Initialize the Database Connection
func getSession(collectionName string) *mgo.Collection {
	log.Println("Initializing Database Connection")
	dbURL := config.Database.DbHost
	mgoSession, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatalf("Error! %v", err)
	}
	session := mgoSession.Clone()
	defer session.Close()
	return session.DB(databaseName).C(collectionName)
}

// ExecuteQuery ..
func ExecuteQuery(collectionName string, s func(*mgo.Collection) error) error {
	session := getSession(collectionName)
	return s(session)
}
