package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/laidingqing/amadd9/common/config"
	"github.com/laidingqing/amadd9/common/registry"
	"github.com/laidingqing/amadd9/common/util"
	"github.com/laidingqing/amadd9/config/config_loader"
	"github.com/laidingqing/amadd9/config/config_service"
	"gopkg.in/natefinch/lumberjack.v2"
)

func parseCmdParams() string {
	defaultConfig, err := util.DefaultConfigLocation()
	if err != nil {
		log.Fatalf("Error setting config file: %v", err)
	}
	hostName := flag.String("hostName", "localhost", "The host name for this instance")
	nodeId := flag.String("nodeId", "cfg1", "The node Id for this instance")
	port := flag.String("port", "4140", "The port number for this instance")
	useSSL := flag.Bool("useSSL", false, "use SSL")
	sslCertFile := flag.String("sslCertFile", "", "The SSL certificate file")
	sslKeyFile := flag.String("sslKeyFile", "", "The SSL key file")
	registryLocation := flag.String("registryLocation", "http://localhost:2379", "URL for etcd")
	configFile := flag.String("config", defaultConfig, "config file to load")
	flag.Parse()
	config.Service.DomainName = *hostName
	config.Service.NodeId = *nodeId
	config.Service.Port = *port
	config.Service.UseSSL = *useSSL
	config.Service.SSLCertFile = *sslCertFile
	config.Service.SSLKeyFile = *sslKeyFile
	config.Service.RegistryLocation = *registryLocation
	return *configFile
}

func main() {
	// Get command line arguments
	configFile := parseCmdParams()
	// Load Configuration
	config.LoadConfig(configFile)
	// Set up Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   config.Logger.LogFile,
		MaxSize:    config.Logger.MaxSize,
		MaxBackups: config.Logger.MaxBackups,
		MaxAge:     config.Logger.MaxAge,
	})
	// Initialize our etcd and couchdb connections
	config_loader.InitRegistry()
	config_loader.InitDatabase()
	// Clear out any old config that may be hanging around
	config_loader.ClearConfig()
	// Set the configuration keys in etcd
	config_loader.SetConfig()
	//Start up the config service
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	//Enable GZIP support
	wsContainer.EnableContentEncoding(true)
	cc := config_service.ConfigController{}
	cc.Register(wsContainer)
	registry.Init("Config", registry.ConfigServiceLocation)
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
