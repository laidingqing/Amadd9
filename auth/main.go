package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/auth/auth_service"
	"github.com/laidingqing/amadd9/common/config"
	"github.com/laidingqing/amadd9/common/database"
	"github.com/laidingqing/amadd9/common/registry"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	config.LoadDefaults()
	config.ParseCmdParams(config.DefaultCmdLine{
		HostName:         "localhost",
		NodeId:           "auth",
		Port:             "4130",
		UseSSL:           false,
		RegistryLocation: "http://localhost:2379", //Etcd endpoints
	})
	config.InitEtcd()
	config.FetchCommonConfig()
	// config.FetchServiceSection(config.UserService)
	// config.FetchServiceSection(config.AuthService)
	log.SetOutput(&lumberjack.Logger{
		Filename:   config.Logger.LogFile,
		MaxSize:    config.Logger.MaxSize,
		MaxBackups: config.Logger.MaxBackups,
		MaxAge:     config.Logger.MaxAge,
	})
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	//Enable Gzip support
	wsContainer.EnableContentEncoding(true)
	uc := auth_service.AuthController{}
	uc.Register(wsContainer)
	database.InitDb()
	registry.Init(registry.AuthEndpointName(), registry.AuthLocation)
	httpAddr := ":" + config.Service.Port
	if config.Service.UseSSL == true {
		certFile := config.Service.SSLCertFile
		keyFile := config.Service.SSLKeyFile
		log.Fatal(http.ListenAndServeTLS(httpAddr,
			certFile, keyFile, wsContainer))
	} else {
		log.Fatal(http.ListenAndServe(httpAddr, wsContainer))
	}
}
