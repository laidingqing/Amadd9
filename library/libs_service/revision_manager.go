package libs_service

import (
	"time"

	. "github.com/laidingqing/amadd9/common/database"
	. "github.com/laidingqing/amadd9/common/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	revisionDbCollection = "revisions"
)

//RevisionManager revision db manager
type RevisionManager struct{}

func (rm *RevisionManager) index(tabID string, rlr *RevisionsListResponse) (string, error) {
	var list []RevisionRecord
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"tabsID": tabID}).All(&list)
	}
	err := ExecuteQuery(revisionDbCollection, query)
	if err != nil {
		return "", err
	}
	rlr.Revisions = list
	return "", nil
}

func (rm *RevisionManager) create(revision *RevisionRecord, rr *RevisionResponse) (string, error) {
	query := func(c *mgo.Collection) error {
		revision.ID = bson.NewObjectId()
		revision.UploadedAt = time.Now()
		return c.Insert(revision)
	}
	err := ExecuteQuery(revisionDbCollection, query)
	if err != nil {
		return "", err
	}

	return revision.ID.Hex(), nil
}
