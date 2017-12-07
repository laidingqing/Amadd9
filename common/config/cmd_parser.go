package config

import (
	"flag"
)

type DefaultCmdLine struct {
	HostName         string
	NodeId           string
	Port             string
	UseSSL           bool
	SSLCertFile      string
	SSLKeyFile       string
	RegistryLocation string
}

func ParseCmdParams(defaults DefaultCmdLine) {
	hostName := flag.String("hostName", defaults.HostName, "The host name for this instance")
	nodeId := flag.String("nodeId", defaults.NodeId, "The node Id for this instance")
	port := flag.String("port", defaults.Port, "The port number for this instance")
	useSSL := flag.Bool("useSSL", defaults.UseSSL, "use SSL")
	sslCertFile := flag.String("sslCertFile", defaults.SSLCertFile, "The SSL certificate file")
	sslKeyFile := flag.String("sslKeyFile", defaults.SSLKeyFile, "The SSL key file")
	registryLocation := flag.String("registryLocation", defaults.RegistryLocation, "URL for etcd")
	flag.Parse()
	Service.DomainName = *hostName
	Service.NodeId = *nodeId
	Service.Port = *port
	Service.UseSSL = *useSSL
	Service.SSLCertFile = *sslCertFile
	Service.SSLKeyFile = *sslKeyFile
	Service.RegistryLocation = *registryLocation
}
