package libs_service

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
	"gopkg.in/mgo.v2/bson"
)

//TabsController tabs版本控制器
type TabsController struct{}

var tabUri = "/tabs"

type TabsResponse struct {
}

// AddRoutes add route to webservice
func (lc TabsController) AddRoutes(ws *restful.WebService) {
	ws.Route(ws.POST(tabUri + "").To(lc.create).
		Doc("create new tabs in library.").
		Operation("postTabs").
		Writes(TabsResponse{}))
}

//create a tab library by artist
func (lc TabsController) create(request *restful.Request, response *restful.Response) {
	curUser := GetCurrentUser(request, response)
	newTab := new(TabRecord)
	err := request.ReadEntity(newTab)
	if err != nil {
		WriteBadRequestError(response)
		return
	}
	am := ArtistManager{}
	artistID, err := am.Create(&newTab.Artist, curUser)
	if err != nil {
		WriteError(err, response)
		return
	}
	newTab.ArtistID = artistID
	rev, _ := new(TabsManager).Create(newTab, curUser)
	response.AddHeader("ETag", rev)
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(&TabRecord{
		ID: bson.ObjectIdHex(rev),
	})
}
