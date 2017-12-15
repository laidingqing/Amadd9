package libs_service

import (
	"time"

	. "github.com/laidingqing/amadd9/common/database"
	. "github.com/laidingqing/amadd9/common/entities"
	slugification "github.com/rhinoman/go-slugification"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	tabsDbCollection = "tabs"
)

// ArtistManager artist db manager
type TabsManager struct{}

// SetUp Do't use , tests only
func (tm *TabsManager) SetUp() (string, error) {
	return "", nil
}

//Create a new artist record.
func (tm *TabsManager) Create(newTab *TabRecord, curUser *CurrentUserInfo) (string, error) {
	query := func(c *mgo.Collection) error {
		newTab.ID = bson.NewObjectId()
		newTab.Slug = slugification.Slugify(newTab.Artist.Name + " " + newTab.Name)
		newTab.CreatedAt = time.Now()
		return c.Insert(newTab)
	}
	err := ExecuteQuery(tabsDbCollection, query)
	if err != nil {
		return "", err
	}
	return newTab.ID.Hex(), nil
}

//Get get a tab record
func (tm *TabsManager) Get(tabID string) (string, error) {
	tr := TabRecord{}
	query := func(c *mgo.Collection) error {
		return c.FindId(bson.ObjectIdHex(tabID)).One(&tr)
	}
	err := ExecuteQuery(tabsDbCollection, query)
	if err != nil {
		return "", err
	}
	return tr.ID.Hex(), nil
}
