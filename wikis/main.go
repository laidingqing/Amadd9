package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/common/config"
	"github.com/laidingqing/amadd9/common/database"
	"github.com/laidingqing/amadd9/common/registry"
	"github.com/laidingqing/amadd9/wikis/wiki_service"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Load default config
	config.LoadDefaults()
	// Parse the command line parameters
	config.ParseCmdParams(config.DefaultCmdLine{
		HostName:         "localhost",
		NodeId:           "ws1",
		Port:             "4110",
		UseSSL:           false,
		RegistryLocation: "http://localhost:2379",
	})
	// Fetch configuration from etcd
	config.InitEtcd()
	config.FetchCommonConfig()
	config.FetchServiceSection(config.WikiService)
	// Set up Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   config.Logger.LogFile,
		MaxSize:    config.Logger.MaxSize,
		MaxBackups: config.Logger.MaxBackups,
		MaxAge:     config.Logger.MaxAge,
	})
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	//Enable Gzip
	wsContainer.EnableContentEncoding(true)
	wc := wiki_service.WikisController{}
	wc.Register(wsContainer)
	database.InitDb()
	registry.Init("Wikis", registry.WikisLocation)
	httpAddr := ":" + config.Service.Port
	if config.Service.UseSSL == true {
		certFile := config.Service.SSLCertFile
		keyFile := config.Service.SSLKeyFile
		log.Fatal(http.ListenAndServeTLS(httpAddr, certFile, keyFile, wsContainer))
	} else {
		log.Fatal(http.ListenAndServe(httpAddr, wsContainer))
	}
}
