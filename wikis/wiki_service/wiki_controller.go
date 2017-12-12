package wiki_service

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
)

type WikisController struct{}

type wikiLinks struct {
	HatLinks
	PageIndex  *HatLink `json:"index,omitempty"`
	Search     *HatLink `json:"search,omitempty"`
	CreatePage *HatLink `json:"create_page,omitempty"`
}

type WikiRecordResponse struct {
	Links      wikiLinks  `json:"_links"`
	WikiRecord WikiRecord `json:"wiki_record"`
}

type WikiIndexResponse struct {
	Links         HatLinks      `json:"_links"`
	TotalRows     int           `json:"totalRows"`
	PageNum       int           `json:"offset"`
	WikiIndexList WikiIndexList `json:"_embedded"`
}

type WikiIndexList struct {
	List []WikiRecordResponse `json:"ea:wiki"`
}

func (wc WikisController) wikiUri() string {
	return ApiPrefix() + "/wikis"
}

var wikisWebService *restful.WebService

func (wc WikisController) Service() *restful.WebService {
	return wikisWebService
}

func (wc WikisController) genWikiUri(wikiId string) string {
	return wc.wikiUri() + "/" + wikiId
}

//Define routes
func (wc WikisController) Register(container *restful.Container) {
	// pc := PagesController{}
	// fc := FileController{}
	wikisWebService = new(restful.WebService)
	wikisWebService.Filter(LogRequest).
		Filter(AuthUser).
		ApiVersion(ApiVersion()).
		Path(wc.wikiUri()).
		Doc("Manage Wikis").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// wikisWebService.Route(wikisWebService.GET("").To(wc.index).
	// 	Doc("Get a list of wikis").
	// 	Operation("index").
	// 	Param(wikisWebService.QueryParameter("pageNum", "Page Number").DataType("integer")).
	// 	Param(wikisWebService.QueryParameter("numPerPage", "Number of records to return").DataType("integer")).
	// 	Param(wikisWebService.QueryParameter("memberOnly", "Only show wikis user belongs to").DataType("boolean")).
	// 	Writes(WikiIndexResponse{}))

	wikisWebService.Route(wikisWebService.POST("").To(wc.create).
		Doc("Create a new wiki").
		Operation("create").
		Reads(WikiRecord{}).
		Writes(WikiRecordResponse{}))
	//
	// wikisWebService.Route(wikisWebService.GET("/{wiki-id}").To(wc.read).
	// 	Doc("Fetch a Wiki Record").
	// 	Operation("read").
	// 	Param(wikisWebService.PathParameter("wiki-id", "Wiki identifier").DataType("string")).
	// 	Writes(WikiRecordResponse{}))
	//
	// wikisWebService.Route(wikisWebService.GET("/slug/{wiki-slug}").To(wc.readBySlug).
	// 	Doc("Fetch a Wiki Record by its slug").
	// 	Operation("readBySlug").
	// 	Param(wikisWebService.PathParameter("wiki-slug", "Wiki Slug").DataType("string")).
	// 	Writes(WikiRecordResponse{}))
	//
	// wikisWebService.Route(wikisWebService.PUT("/{wiki-id}").To(wc.update).
	// 	Doc("Update a Wiki Record").
	// 	Operation("update").
	// 	Param(wikisWebService.PathParameter("wiki-id", "Wiki identifier").DataType("string")).
	// 	Param(wikisWebService.HeaderParameter("If-Match", "Revision").DataType("string")).
	// 	Reads(WikiRecord{}).
	// 	Writes(WikiRecordResponse{}))
	//
	// wikisWebService.Route(wikisWebService.DELETE("/{wiki-id}").To(wc.del).
	// 	Doc("Delete a Wiki").
	// 	Operation("del").
	// 	Param(wikisWebService.PathParameter("wiki-id", "Wiki identifier").DataType("string")).
	// 	Writes(BooleanResponse{}))

	// pc.AddRoutes(wikisWebService)
	// fc.AddRoutes(wikisWebService)
	container.Add(wikisWebService)
}

//create Create a new wiki
func (wc WikisController) create(request *restful.Request, response *restful.Response) {
	curUser := GetCurrentUser(request, response)
	if curUser == nil {
		Unauthenticated(request, response)
		return
	}
	theWiki := new(WikiRecord)
	err := request.ReadEntity(theWiki)
	if err != nil {
		WriteBadRequestError(response)
		return
	}

	// response.AddHeader("ETag", rev)
	response.WriteHeader(http.StatusCreated)
	// response.WriteEntity(wr)
}
