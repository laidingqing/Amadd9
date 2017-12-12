package wiki_service

import (
	"net/url"
	"time"

	. "github.com/laidingqing/amadd9/common/entities"
	"github.com/rhinoman/go-slugification"
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

type WikiManager struct{}

func WikiDbName(id string) string {
	return "wiki_" + id
}

//Create a new wiki
func (wm *WikiManager) Create(id string, wr *WikiRecord, curUser *CurrentUserInfo) (string, error) {
	wr.ID = bson.NewObjectId()
	wr.Slug = slugification.Slugify(wr.Name)
	wr.CreatedAt = time.Now().UTC()
	wr.ModifiedAt = time.Now().UTC()
	wr.Type = "wiki_record"

	return "", nil
}

// checkForDuplicateSlug Check for duplicate wiki slug
func (wm *WikiManager) checkForDuplicateSlug(slug string) error {
	params := url.Values{}
	params.Add("key", "\""+slug+"\"")
	params.Add("group", "true")
	// response := CheckSlugResponse{}

	return nil
}
