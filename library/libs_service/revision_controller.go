package libs_service

import restful "github.com/emicklei/go-restful"

//RevisionController tabs版本控制器
type RevisionController struct{}

var revisionUri = "/{tabs-id}/revisions"

func (lc LibraryController) AddRoutes(ws *restful.WebService) {

}
