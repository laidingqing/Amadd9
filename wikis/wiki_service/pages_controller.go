package wiki_service

import (
	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
	"github.com/laidingqing/wikifeat/wikis/wiki_service/wikit"
)

var pageUri = "/{wiki-id}/pages"

type PagesController struct{}

type PageIndexResponse struct {
	Links         HatLinks      `json:"_links"`
	PageIndexList PageIndexList `json:"_embedded"`
}

type PageIndexList struct {
	List []PageIndexItem `json:"ea:page"`
}

type PageIndexItem struct {
	Entry wikit.PageIndexEntry `json:"page"`
}

func (pc PagesController) AddRoutes(ws *restful.WebService) {
	ws.Route(ws.GET(pageUri).To(pc.index).
		Doc("Get list of pages in this wiki").
		Operation("index").
		Param(ws.PathParameter("wiki-id", "Wiki identifier").DataType("string")).
		Writes(PageIndexResponse{}))
}

//Get Page index
func (pc PagesController) index(request *restful.Request, response *restful.Response) {
	// curUser := GetCurrentUser(request, response)
	// if curUser == nil {
	// 	Unauthenticated(request, response)
	// 	return
	// }
	wikiId := request.PathParameter("wiki-id")
	pIndex, err := new(PageManager).Index(wikiId, curUser)
	if err != nil {
		WriteError(err, response)
		return
	}
	indexResponse := pc.getIndexResponse(wikiId, curUser, pIndex)
	wikiUri := ApiPrefix() + "/wikis/" + wikiId
	indexResponse.Links = GenIndexLinks(curUser.User.Roles,
		"wiki_"+wikiId, wikiUri)
	SetAuth(response, curUser.Auth)
	response.WriteEntity(indexResponse)
}
