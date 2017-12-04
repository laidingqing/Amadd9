package services

import (
	"log"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/common/config"
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
