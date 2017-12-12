package libs_service

import (
	restful "github.com/emicklei/go-restful"
)

//TabsController tabs版本控制器
type TabsController struct{}

var tabUri = "/tabs"

type TabsResponse struct {
}

func (lc TabsController) AddRoutes(ws *restful.WebService) {
	ws.Route(ws.POST(tabUri + "").To(lc.create).
		Doc("create new tabs in library.").
		Operation("postTabs").
		Writes(TabsResponse{}))
}

func (lc TabsController) create(request *restful.Request, response *restful.Response) {

}
