package user_service

import (
	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/database"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
)

type UsersController struct{}

func (uc UsersController) userUri() string {
	return ApiPrefix() + "/users"
}

var usersWebService *restful.WebService

func (uc UsersController) Service() *restful.WebService {
	return usersWebService
}

func (uc UsersController) Register(container *restful.Container) {
	usersWebService = new(restful.WebService)
	usersWebService.Filter(LogRequest)
	usersWebService.
		Path(uc.userUri()).
		Doc("Manage Users").
		ApiVersion(ApiVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

}
