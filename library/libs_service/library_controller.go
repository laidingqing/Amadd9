package libs_service

import (
	"log"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/services"
)

type LibraryController struct{}

func (uc LibraryController) libsUri() string {
	return ApiPrefix() + "/libs"
}

var libsWebService *restful.WebService

// Service , LibraryController Service
func (uc LibraryController) Service() *restful.WebService {
	return libsWebService
}

// Register register webservice
func (uc LibraryController) Register(container *restful.Container) {

	lc := LibraryController{}
	ac := ArtistController{}
	tc := TabsController{}

	log.Println("Register restful.Container")
	libsWebService = new(restful.WebService)
	libsWebService.Filter(LogRequest)
	libsWebService.
		Path(uc.libsUri()).
		Doc("Manage Libs: Tabs, Artist, Annotation").
		ApiVersion(ApiVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tc.AddRoutes(libsWebService)
	lc.AddRoutes(libsWebService)
	ac.AddRoutes(libsWebService)
	container.Add(libsWebService)
}
