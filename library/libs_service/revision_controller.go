package libs_service

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
)

var revisionURI = "/{tabs-id}/revisions"

//RevisionController tabs版本控制器
type RevisionController struct{}

//RevisionsListResponse list revisions response
type RevisionsListResponse struct {
	Revisions []RevisionRecord `json:"revisions"`
}

type RevisionResponse struct {
	Revision RevisionRecord `json:"revision"`
}

//AddRoutes add revision routes to Webservice
func (rc RevisionController) AddRoutes(ws *restful.WebService) {
	ws.Route(ws.GET(revisionURI).To(rc.index).
		Filter(FindByTabFilter).
		Doc("Get list of revision in this tabs").
		Operation("index").
		Param(ws.PathParameter("tabs-id", "Tab identifier").DataType("string")).
		Writes(RevisionsListResponse{}))

	ws.Route(ws.POST(revisionURI).To(rc.create).
		Filter(FindByTabFilter).
		Doc("create a revision in this tabs").
		Operation("create").
		Param(ws.PathParameter("tabs-id", "Tab identifier").DataType("string")).
		Writes(RevisionResponse{}))
}

// private func index, find all revision by tab-id
func (rc RevisionController) index(request *restful.Request, response *restful.Response) {
	tabID := request.PathParameter("tabs-id")
	rm := new(RevisionManager)
	rlr := RevisionsListResponse{}
	_, err := rm.index(tabID, &rlr)
	if err != nil {
		WriteError(err, response)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.WriteEntity(rlr)
}

// private func create, create a revision by tab-id
func (rc RevisionController) create(request *restful.Request, response *restful.Response) {
	tabID := request.PathParameter("tabs-id")
	rm := new(RevisionManager)
	rr := RevisionResponse{}
	//TODO file uploader
	newRevision := new(RevisionRecord)
	newRevision.TabsID = tabID
	err := request.ReadEntity(newRevision)
	if err != nil {
		WriteError(err, response)
		return
	}
	rev, err := rm.create(newRevision, &rr)
	if err != nil {
		WriteError(err, response)
		return
	}
	log.Printf("revision created, ID:%v", rev)
	response.WriteHeader(http.StatusOK)
	response.WriteEntity(&rr)
}

//FindByTabFilter ..jwt token validate.
func FindByTabFilter(request *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	tm := new(TabsManager)
	_, err := tm.Get(request.PathParameter("tabs-id"))
	if err != nil {
		NoFoundTabRecord(request, resp)
		return
	}
	chain.ProcessFilter(request, resp)
}
