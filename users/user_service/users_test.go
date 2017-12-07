package user_service_test

import (
	"testing"

	"github.com/laidingqing/amadd9/common/config"
)

func setup() {
	config.LoadDefaults()
}

func TestUsers(t *testing.T) {
	setup()
	config.LoadDefaults()
	//Parse the command line parameters
	config.ParseCmdParams(config.DefaultCmdLine{
		HostName:         "localhost",
		NodeId:           "test1",
		Port:             "4130",
		UseSSL:           false,
		RegistryLocation: "http://localhost:2379",
	})
	config.InitEtcd()
	config.FetchCommonConfig()
}
