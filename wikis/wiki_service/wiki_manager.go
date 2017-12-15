package wiki_service

import (
	"net/url"
	"time"

	. "github.com/laidingqing/amadd9/common/database"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/util"
	"github.com/rhinoman/go-slugification"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WikiRecordListResult struct {
	Id    string     `json:"id,omitempty"`
	Key   string     `json:"key"`
	Value WikiRecord `json:"value"`
}

type WikiListResponse struct {
	ViewResponse
	Rows []WikiRecordListResult `json:"rows,omitempty"`
}

type WikiSlugViewResponse struct {
	ViewResponse
	Rows []WikiSlugViewResult `json:"rows,omitempty"`
}

type WikiSlugViewResult struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value struct {
		Rev        string     `json:"wikiRev"`
		WikiRecord WikiRecord `json:"wiki_record"`
	} `json:"value"`
}

type CheckSlugResponse struct {
	Rows []KvItem `json:"rows"`
}

type KvItem struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

//WikiManager wiki database manager
type WikiManager struct{}

var (
	wikiDbCollection = "wikis"
)

//Create a new wiki
func (wm *WikiManager) Create(wr *WikiRecord, curUser *CurrentUserInfo) (string, error) {
	py := new(Pinyin)
	wr.ID = bson.NewObjectId()
	p, _ := py.Convert(wr.Name)

	wr.Slug = slugification.Slugify(p)
	wr.CreatedAt = time.Now().UTC()
	wr.ModifiedAt = time.Now().UTC()
	wr.Type = "wiki_record"

	query := func(c *mgo.Collection) error {
		return c.Insert(wr)
	}

	err := ExecuteQuery(wikiDbCollection, query)
	if err != nil {
		return "", err
	}

	return wr.ID.Hex(), nil
}

//ReadBySlug Fetch a wiki record by its slug
func (wm *WikiManager) ReadBySlug(slug string, wikiRecord *WikiRecord, curUser *CurrentUserInfo) (string, error) {

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"slug": slug}).One(&wikiRecord)
	}
	err := ExecuteQuery(wikiDbCollection, query)
	if err != nil {
		return "", NotFoundError()
	}
	return wikiRecord.ID.Hex(), nil
}

// checkForDuplicateSlug Check for duplicate wiki slug
func (wm *WikiManager) checkForDuplicateSlug(slug string) error {
	params := url.Values{}
	params.Add("key", "\""+slug+"\"")
	params.Add("group", "true")
	// response := CheckSlugResponse{}

	return nil
}
