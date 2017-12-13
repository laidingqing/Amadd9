package libs_service

import (
	. "github.com/laidingqing/amadd9/common/database"
	. "github.com/laidingqing/amadd9/common/entities"
	couchdb "github.com/rhinoman/couchdb-go"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	artistDbCollection = "artists"
)

// ArtistManager artist db manager
type ArtistManager struct{}

// SetUp Do't use , tests only
func (am *ArtistManager) SetUp() (string, error) {
	return "", nil
}

//Create a new artist record.
func (am *ArtistManager) Create(newArtist *ArtistRecord, curUser *CurrentUserInfo) (string, error) {
	if bson.IsObjectIdHex(newArtist.ID.Hex()) {
		query := func(c *mgo.Collection) error {
			return c.FindId(newArtist.ID).One(newArtist)
		}
		_ = ExecuteQuery(artistDbCollection, query)
		return "", nil
	}
	query := func(c *mgo.Collection) error {
		newArtist.ID = bson.NewObjectId()
		return c.Insert(newArtist)
	}
	id, _ := am.findByName(newArtist.Name)
	if id != "" {
		return id, nil
	}
	err := ExecuteQuery(artistDbCollection, query)
	if err != nil {
		return "", err
	}
	return newArtist.ID.Hex(), nil
}

func (am *ArtistManager) findByName(name string) (string, error) {
	var artist ArtistRecord
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"name": name}).One(&artist)
	}
	err := ExecuteQuery(artistDbCollection, query)
	if err != nil {
		return "", &couchdb.Error{
			StatusCode: 500,
			Reason:     "Db error :" + err.Error(),
		}
	}
	return artist.ID.Hex(), nil
}
