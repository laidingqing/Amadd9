package config_service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/laidingqing/amadd9/common/config"
	"github.com/laidingqing/amadd9/common/registry"
)

type ConfigManager struct{}

var configLocation = registry.EtcdPrefix + "/config/"

//Get a config parameter
func (cm *ConfigManager) getConfigParam(section string,
	paramName string) (string, error) {
	serviceSection, err := config.ServiceSectionFromString(section)
	if err != nil {
		return "", err
	}
	switch serviceSection {
	case config.AuthService:
		return cm.getAuthParam(strings.ToLower(paramName))
	default:
		return "", errors.New("Invalid config section requested")
	}

}

//Get an Auth config parameter
func (cm *ConfigManager) getAuthParam(paramName string) (string, error) {
	switch paramName {
	case "allowguestaccess":
		return strconv.FormatBool(config.Auth.AllowGuest), nil
	case "allownewuserregistration":
		return strconv.FormatBool(config.Auth.AllowNewUserRegistration), nil
	default:
		return "", errors.New("Invalid auth config parameter requested")
	}
}
