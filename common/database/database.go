package database

import (
	"log"

	"github.com/laidingqing/amadd9/common/config"
	couchdb "github.com/rhinoman/couchdb-go"
	mgo "gopkg.in/mgo.v2"
)

var (
	mgoSession   *mgo.Session
	databaseName string
)

func InitDb() {
	databaseName = config.Database.MainDb
}

func SetupDb() {

}

func NotAdminError() error {
	return &couchdb.Error{
		StatusCode: 403,
		Reason:     "Not an admin",
	}
}

// ExecuteQuery ..
func ExecuteQuery(collectionName string, s func(*mgo.Collection) error) error {
	dbURL := config.Database.DbHost
	mgoSession, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatalf("Error! %v", err)
	}
	session := mgoSession.Clone()
	defer session.Close()
	return s(session.DB(databaseName).C(collectionName))
}

// ExecuteCount count query
func ExecuteCount(collectionName string, s func(*mgo.Collection) (int, error)) (int, error) {
	dbURL := config.Database.DbHost
	mgoSession, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatalf("Error! %v", err)
	}
	session := mgoSession.Clone()
	defer session.Close()
	return s(session.DB(databaseName).C(collectionName))
}
