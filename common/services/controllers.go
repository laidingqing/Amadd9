package services

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/auth"
	"github.com/laidingqing/amadd9/common/config"
	. "github.com/laidingqing/amadd9/common/entities"
	couchdb "github.com/rhinoman/couchdb-go"
)

// Controller ..
type Controller interface {
	Service() *restful.WebService
}

// HatLinks ..
type HatLinks struct {
	Self   *HatLink `json:"self"`
	Create *HatLink `json:"create,omitempty"`
	Update *HatLink `json:"update,omitempty"`
	Delete *HatLink `json:"delete,omitempty"`
}

// HatLink ..
type HatLink struct {
	Href       string `json:"href"`
	HrefLang   string `json:"hreflang,omitempty"`
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	Deprecated string `json:"deprecation,omitempty"`
	Name       string `json:"name,omitempty"`
	Profile    string `json:"profile,omitempty"`
	Templated  bool   `json:"templated,omitempty"`
	Method     string `json:"method,omitempty"` //HTTP method, not standard HAL
}

// BooleanResponse ..
type BooleanResponse struct {
	Success bool `json:"success"`
}

// ApiVersion ..
func ApiVersion() string {
	return config.ApiVersion
}

// ApiPrefix ...
func ApiPrefix() string {
	return "/api/" + ApiVersion()
}

//LogRequest Filter function.  Logs incoming requests
func LogRequest(request *restful.Request, resp *restful.Response,
	chain *restful.FilterChain) {
	method := request.Request.Method
	url := request.Request.URL.String()
	remoteAddr := request.Request.RemoteAddr
	log.Printf("[API] %v : %v %v", remoteAddr, method, url)
	chain.ProcessFilter(request, resp)
}

// LogError Filter function . Log an error
func LogError(request *restful.Request, resp *restful.Response, err error) {
	method := request.Request.Method
	url := request.Request.URL.String()
	remoteAddr := request.Request.RemoteAddr
	log.Printf("[ERROR] %v : %v : %v %v", err, remoteAddr, method, url)
}

// WriteBadRequestError write err to response
func WriteBadRequestError(response *restful.Response) {
	log.Printf("400: Bad Request")
	response.WriteErrorString(http.StatusBadRequest, "Bad Request")
}

func WriteIllegalRequestError(response *restful.Response) {
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusBadRequest, "Bad Request")
}

//WriteServerError , Writes an internal server error
func WriteServerError(err error, response *restful.Response) {
	log.Printf("%v", err)
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusInternalServerError, err.Error())
}

//Writes and logs errors from the couchdb driver
func WriteError(err error, response *restful.Response) {
	var statusCode int
	var reason string = "error"
	//Is this a couchdb error?
	cErr, ok := err.(*couchdb.Error)
	if ok { // Yes!
		statusCode = cErr.StatusCode
		reason = cErr.Reason
	} else { // No, try to parse :(
		str := err.Error()
		errStrings := strings.Split(str, ":")
		statusCode := 0
		var cErr error
		if len(errStrings) > 1 {
			statusCode, cErr = strconv.Atoi(errStrings[1])
			reason = http.StatusText(statusCode)
		}
		if cErr != nil || statusCode == 0 {
			statusCode = 500
		}
	}
	//Write the error to the response
	response.WriteErrorString(statusCode, reason)
	//Log the error
	log.Printf("%v", err)
}

//GetAdminUser , Returns the Admin Credentials as a CurrentUserInfo
func GetAdminUser() *CurrentUserInfo {
	return &CurrentUserInfo{
		Roles: []string{"admin"},
		User: &User{
			Roles: []string{"admin"},
		},
	}
}

// GetCurrentUser , Get current session user.
func GetCurrentUser(request *restful.Request, response *restful.Response) *CurrentUserInfo {
	curUser, ok := request.Attribute("currentUser").(*CurrentUserInfo)
	if ok == false || curUser == nil {
		return nil
	}
	return curUser
}

//Unauthenticated , Writes unauthenticated error to response
func Unauthenticated(request *restful.Request, response *restful.Response) {
	LogError(request, response, errors.New("Unauthenticated"))
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(401, "Unauthenticated")
}

//AuthUser ..jwt token validate.
func AuthUser(request *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	res := ValidateToken(request.Request)
	if !res {
		Unauthenticated(request, resp)
		return
	}
	chain.ProcessFilter(request, resp)
}
