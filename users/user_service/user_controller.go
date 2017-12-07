package user_service

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/common/config"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
)

type UsersController struct{}

func (uc UsersController) userUri() string {
	return ApiPrefix() + "/users"
}

var usersWebService *restful.WebService

// Service , UsersController Service
func (uc UsersController) Service() *restful.WebService {
	return usersWebService
}

// Register register webservice
func (uc UsersController) Register(container *restful.Container) {
	usersWebService = new(restful.WebService)
	usersWebService.Filter(LogRequest)
	usersWebService.
		Path(uc.userUri()).
		Doc("Manage Users").
		ApiVersion(ApiVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	usersWebService.Route(usersWebService.POST("/register").To(uc.register).
		Doc("New user Registration").
		Operation("register").
		Reads(User{}))
}

//New user (self) registration
func (uc UsersController) register(request *restful.Request, response *restful.Response) {
	if config.Auth.AllowNewUserRegistration {
		newUser := new(User)
		if err := request.ReadEntity(newUser); err != nil {
			WriteBadRequestError(response)
			return
		}
		rev, err := new(UserManager).Register(newUser)
		if err != nil {
			WriteError(err, response)
			return
		}
		response.AddHeader("ETag", rev)
		response.WriteHeader(http.StatusCreated)
	} else {
		response.WriteHeader(http.StatusForbidden)
	}
}
