package config

import (
	"log"
	"path"

	"github.com/laidingqing/amadd9/common/util"
)

var ApiVersion = "v1"

var Service struct {
	DomainName       string
	NodeId           string
	Port             string
	ApiVersion       string
	RegistryLocation string
	UseSSL           bool
	SSLCertFile      string
	SSLKeyFile       string
}

var Frontend struct {
	WebAppDir string
	PluginDir string
	Homepage  string
}

var Search struct {
	SearchServerLocation string
}

var Database struct {
	DbHost          string
	DbAdminUser     string
	DbAdminPassword string
	DbTimeout       string
	MainDb          string
}

var Logger struct {
	LogFile    string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

var Auth struct {
	Authenticator            string
	SessionTimeout           uint64
	PersistentSessions       bool
	AllowGuest               bool
	AllowNewUserRegistration bool
	MinPasswordLength        int
}

var ServiceRegistry struct {
	EntryTTL             uint64
	CacheRefreshInterval uint64
}

var Users struct {
	AvatarDb string
}

// LoadDefaults ..Initialize Default values
func LoadDefaults() {
	execDir, err := util.GetExecDirectory()
	if err != nil {
		log.Fatal(err)
	}
	ServiceRegistry.CacheRefreshInterval = 75
	ServiceRegistry.EntryTTL = 60
	Frontend.WebAppDir = path.Join(execDir, "web_app/app")
	Frontend.PluginDir = path.Join(execDir, "plugins")
	Frontend.Homepage = ""
	Database.DbHost = "localhost"
	Database.DbAdminUser = "adminuser"
	Database.DbAdminPassword = "password"
	Database.DbTimeout = "0"
	Database.MainDb = "main_ut"
	Logger.LogFile = "amadd9-service.log"
	Logger.MaxSize = 10
	Logger.MaxBackups = 3
	Logger.MaxAge = 30
	Auth.Authenticator = "standard"
	Auth.SessionTimeout = 600
	Auth.PersistentSessions = true
	Auth.AllowGuest = true
	Auth.AllowNewUserRegistration = false
	Auth.MinPasswordLength = 6
	Users.AvatarDb = "avatar_ut"
}
