package config_service

import (
	"github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/services"
)

/**
  We need to be able to query a few select configuration parameters at runtime.
  Might eventually expand this to allow on-the-fly configuration changes.
*/

type ConfigController struct{}

type ConfigResponse struct {
	Links HatLinks    `json:"_links"`
	Param ConfigParam `json:"configParam"`
}

type ConfigParam struct {
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
}

var configWebService *restful.WebService

func (cc ConfigController) configUri() string {
	return ApiPrefix() + "/config"
}

func (cc ConfigController) Service() *restful.WebService {
	return configWebService
}

//Define routes
func (cc ConfigController) Register(container *restful.Container) {
	configWebService = new(restful.WebService)
	configWebService.Filter(LogRequest)
	configWebService.
		Path(cc.configUri()).
		Doc("Query system configuration").
		ApiVersion(ApiVersion()).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	configWebService.Route(configWebService.GET("/{section}/{parameter}").To(cc.getConfigParam).
		Doc("Get a single configuration parameter value").
		Operation("getConfigParam").
		Param(configWebService.PathParameter("section", "Section").DataType("string")).
		Param(configWebService.PathParameter("parameter", "Parameter").DataType("string")).
		Writes(ConfigResponse{}))

	container.Add(configWebService)
}

// Get an individual config parameter
func (cc ConfigController) getConfigParam(request *restful.Request,
	response *restful.Response) {
	section := request.PathParameter("section")
	param := request.PathParameter("parameter")
	value, err := new(ConfigManager).getConfigParam(section, param)
	if err != nil {
		WriteIllegalRequestError(response)
		return
	}
	configResponse := cc.genConfigResponse(param, value)
	response.WriteEntity(configResponse)
}

func (cc ConfigController) genConfigResponse(paramName, paramValue string) ConfigResponse {
	links := HatLinks{}
	uri := cc.configUri() + "/" + paramName
	links.Self = &HatLink{Href: uri, Method: "GET"}
	return ConfigResponse{
		Links: links,
		Param: ConfigParam{
			ParamName:  paramName,
			ParamValue: paramValue,
		},
	}
}
