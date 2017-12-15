package libs_service

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/laidingqing/amadd9/common/database"
	"gopkg.in/mgo.v2/bson"
)

type FileManager struct{}

//SaveToGridFs save file byte to mongodb fs implements.
func (fm *FileManager) SaveToGridFs(md5 string, body []byte) (string, error) {
	return ExecuteGridFS(bson.NewObjectId().Hex(), md5, body)
}

//Download a file from url
func (fm *FileManager) Download(url string) (string, error) {
	if url == "" {
		return "", nil
	}
	res, err := http.Get(url)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	log.Printf("Download file from: %v", url)
	if err != nil {
		log.Printf("Download file error: %v", err.Error())
		return "", nil
	}
	md5 := fmt.Sprintf("%x", md5.Sum(body))
	return fm.SaveToGridFs(md5, body)
}
