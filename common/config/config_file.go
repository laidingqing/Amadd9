package config

import (
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/alyu/configparser"
	"github.com/laidingqing/amadd9/common/util"
)

// Load config values from file
func LoadConfig(filename string) {
	LoadDefaults()
	log.Printf("\nLoading Configuration from %v\n", filename)
	config, err := configparser.Read(filename)
	if err != nil {
		log.Fatal(err)
	}
	dbSection, err := config.Section("Database")
	if err != nil {
		log.Fatal(err)
	}
	logSection, err := config.Section("Logging")
	if err != nil {
		log.Fatal(err)
	}
	authSection, err := config.Section("Auth")
	if err != nil {
		log.Fatal(err)
	}
	registrySection, err := config.Section("ServiceRegistry")
	if err != nil {
		log.Fatal(err)
	}
	//Optional sections
	frontendSection, err := config.Section("Frontend")
	userSection, err := config.Section("Users")
	searchSection, err := config.Section("Search")
	if frontendSection != nil {
		SetFrontendConfig(frontendSection)
	}
	if searchSection != nil {
		SetSearchConfig(searchSection)
	}
	if userSection != nil {
		setUsersConfig(userSection)
	}
	setDbConfig(dbSection)
	setLogConfig(logSection)
	setAuthConfig(authSection)
	setRegistryConfig(registrySection)
}

// Load Frontend configuration options
func SetFrontendConfig(frontendSection *configparser.Section) {
	execDir, _ := util.GetExecDirectory()
	for key, value := range frontendSection.Options() {
		switch key {
		case "webAppDir":
			if value[0] != '/' {
				Frontend.WebAppDir = path.Join(execDir, value)
			} else {
				Frontend.WebAppDir = value
			}
		case "pluginDir":
			if value[0] != '/' {
				Frontend.PluginDir = path.Join(execDir, value)
			} else {
				Frontend.PluginDir = value
			}
		case "homepage":
			Frontend.Homepage = value
		}
	}
}

// Load Search configuration options
func SetSearchConfig(searchSection *configparser.Section) {
	for key, value := range searchSection.Options() {
		switch key {
		case "searchServerLocation":
			Search.SearchServerLocation = value
		}
	}
}

// Load Database configuration options
func setDbConfig(dbSection *configparser.Section) {
	// for key, value := range dbSection.Options() {
	// 	switch key {
	//
	// 	}
	// }
}

func setIntVal(str string, to *int) {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	} else {
		*to = i
	}
}

func setUint64Val(str string, to *uint64) {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	} else {
		*to = i
	}
}

func stringToBool(str string) bool {
	if str == "true" {
		return true
	} else {
		return false
	}
}

// Load Logging configuration
func setLogConfig(logSection *configparser.Section) {
	for key, value := range logSection.Options() {
		switch key {
		case "logFile":
			Logger.LogFile = value
		case "maxSize":
			setIntVal(value, &Logger.MaxSize)
		case "maxBackups":
			setIntVal(value, &Logger.MaxBackups)
		case "maxAge":
			setIntVal(value, &Logger.MaxAge)
		}
	}
}

// Load Auth configuration
func setAuthConfig(authSection *configparser.Section) {
	for key, value := range authSection.Options() {
		switch key {
		case "authenticator":
			Auth.Authenticator = strings.ToLower(value)
		case "sessionTimeout":
			setUint64Val(value, &Auth.SessionTimeout)
		case "persistentSessions":
			Auth.PersistentSessions = stringToBool(value)
		case "allowGuestAccess":
			Auth.AllowGuest = stringToBool(value)
		case "allowNewUserRegistration":
			Auth.AllowNewUserRegistration = stringToBool(value)
		case "minPasswordLength":
			setIntVal(value, &Auth.MinPasswordLength)
		}
	}
}

// Load Registry configuration
func setRegistryConfig(registrySection *configparser.Section) {
	for key, value := range registrySection.Options() {
		switch key {
		case "entryTTL":
			setUint64Val(value, &ServiceRegistry.EntryTTL)
		case "cacheRefreshInterval":
			setUint64Val(value, &ServiceRegistry.CacheRefreshInterval)
		}
	}
}

// Load Users configuraiton
func setUsersConfig(userSection *configparser.Section) {
	for key, value := range userSection.Options() {
		switch key {
		case "avatarDB":
			Users.AvatarDb = value
		}
	}
}
