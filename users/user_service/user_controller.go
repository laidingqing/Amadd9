package user_service

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/common/config"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
	. "github.com/laidingqing/amadd9/common/util"
)

type UsersController struct{}

type UserResponse struct {
	Links HatLinks `json:"_links"`
	User  User     `json:"user"`
}

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
	log.Println("Register restful.Container")
	usersWebService = new(restful.WebService)
	usersWebService.Filter(LogRequest)
	usersWebService.
		Path(uc.userUri()).
		Doc("Manage Users").
		ApiVersion(ApiVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	usersWebService.Route(usersWebService.POST("").To(uc.create).
		// Filter(AuthUser).
		Doc("Create a User").
		Operation("create").
		Reads(User{}).
		Writes(UserResponse{}))

	usersWebService.Route(usersWebService.POST("/register").To(uc.register).
		Doc("New user Registration").
		Operation("register").
		Reads(User{}))

	usersWebService.Route(usersWebService.POST("/sessions").To(uc.sessions).
		Doc("Login Session").
		Operation("sessions").
		Reads(User{}))

	usersWebService.Route(usersWebService.DELETE("/sessions").To(uc.destroy).
		Doc("Destroy Session").
		Operation("sessions").
		Reads(User{}))

	usersWebService.Route(usersWebService.GET("/{user-id}").To(uc.read).
		Doc("Gets a User").
		Operation("read").
		Param(usersWebService.PathParameter("user-id", "User Name").DataType("string")).
		Writes(UserResponse{}))

	container.Add(usersWebService)
}

// Get a User
func (uc UsersController) read(request *restful.Request, response *restful.Response) {
	log.Println("fetch user")
	response.WriteEntity(UserResponse{
		User: User{
			UserName: "laidingqing",
		},
	})
}

//Create a User
func (uc UsersController) create(request *restful.Request,
	response *restful.Response) {
	curUser := GetCurrentUser(request, response)
	if curUser == nil {
		Unauthenticated(request, response)
		return
	}
	newUser := new(User)
	err := request.ReadEntity(newUser)
	if err != nil {
		WriteBadRequestError(response)
		return
	}
	rev, err := new(UserManager).Create(newUser, curUser)
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", rev)
	response.WriteHeader(http.StatusCreated)
}

//New user (self) registration
func (uc UsersController) register(request *restful.Request, response *restful.Response) {
	if config.Auth.AllowNewUserRegistration {
		newUser := new(User)
		newUser.Slat = string(Krand(5, 3))
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

func (uc UsersController) sessions(request *restful.Request, response *restful.Response) {
	theUser := new(User)
	err := request.ReadEntity(theUser)
	if err != nil {
		WriteBadRequestError(response)
		return
	}

	//TODO check username and password, return a session struct

	response.WriteHeader(http.StatusOK)
	// response.WriteEntity(ur)
}

func (uc UsersController) destroy(request *restful.Request, response *restful.Response) {
	response.WriteHeader(http.StatusOK)
	// response.WriteEntity(ur)
}
