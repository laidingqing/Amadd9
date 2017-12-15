package database

import (
	"log"

	"github.com/laidingqing/amadd9/common/config"
	couchdb "github.com/rhinoman/couchdb-go"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

//NotFoundError data not found
func NotFoundError() error {
	return &couchdb.Error{
		StatusCode: 404,
		Reason:     "Not Found",
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

//ExecuteGridFS execute save file to grid fs.
func ExecuteGridFS(id string, name string, body []byte) (string, error) {
	result := make(map[string]interface{})
	q := bson.M{"filename": name}

	dbURL := config.Database.DbHost
	mgoSession, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatalf("Error! %v", err)
	}
	session := mgoSession.Clone()
	defer session.Close()
	gridfs := session.DB(databaseName).GridFS("fs")
	if err = gridfs.Find(q).One(&result); err == nil {
		return result["_id"].(bson.ObjectId).Hex(), nil
	}
	fs, err := gridfs.Create(id)
	if err != nil {
		return "", err
	}
	defer fs.Close()
	fs.SetName(name)
	if _, err := fs.Write(body); err != nil {
		return "", err
	}
	return name, nil
}
