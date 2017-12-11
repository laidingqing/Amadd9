package auth_service

import (
	"github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/common/auth"
	. "github.com/laidingqing/amadd9/common/services"
)

type AuthController struct{}

var authWebService *restful.WebService

func (ac AuthController) authUri() string {
	return ApiPrefix() + "/auth"
}

func (ac AuthController) Service() *restful.WebService {
	return authWebService
}

//Define routes
func (ac AuthController) Register(container *restful.Container) {
	authWebService = new(restful.WebService)
	authWebService.Filter(LogRequest)
	authWebService.
		Path(ac.authUri()).
		Doc("Manage Authentication").
		ApiVersion(ApiVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	authWebService.Route(authWebService.GET("/token").To(ac.validateAuth).
		Doc("Validation Token").
		Operation("getAuth").
		Writes(BooleanResponse{}))

	authWebService.Route(authWebService.POST("").To(ac.create).
		Doc("Generator a Jwt Token").
		Operation("create").
		Writes(auth.JwtSession{}))

	container.Add(authWebService)

}

// validate jwt session
func (ac AuthController) validateAuth(r *restful.Request, response *restful.Response) {
	valid := auth.ValidateToken(r.Request)
	response.WriteEntity(BooleanResponse{
		Success: valid,
	})
}

// validate jwt session
func (ac AuthController) create(r *restful.Request, response *restful.Response) {
	jwt, err := auth.CreateJWT()
	if err != nil {
		WriteError(err, response)
		return
	}
	response.WriteEntity(auth.JwtSession{
		Token: jwt,
	})
}
