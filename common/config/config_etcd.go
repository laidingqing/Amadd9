package config

import (
	"errors"
	"log"
	"reflect"
	"strconv"

	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// Config Locations in Etcd
var ConfigPrefix = "/amadd9/config/"
var DbConfigLocation = ConfigPrefix + "db/"
var LogConfigLocation = ConfigPrefix + "log/"
var AuthConfigLocation = ConfigPrefix + "auth/"
var UsersConfigLocation = ConfigPrefix + "users/"
var RegistryConfigLocation = ConfigPrefix + "registry/"

// The Etcd keys client
var kapi etcd.KeysAPI

// ServiceSection Service section enum
type ServiceSection int

const (
	NONE ServiceSection = iota
	AuthService
	UserService
	WikiService
	FrontendService
)

func ServiceSectionFromString(sectionStr string) (ServiceSection, error) {
	switch sectionStr {
	case "auth":
		return AuthService, nil
	case "user":
		return UserService, nil
	case "wiki":
		return WikiService, nil
	case "frontend":
		return FrontendService, nil
	default:
		return NONE, errors.New("Unknown section")
	}
}

// InitEtcd init etcd .
func InitEtcd() {
	log.Printf("Initializing etcd config connection")
	etcdCfg := etcd.Config{
		Endpoints: []string{Service.RegistryLocation},
		Transport: etcd.DefaultTransport,
	}
	cli, err := etcd.New(etcdCfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	kapi = etcd.NewKeysAPI(cli)
}

// FetchCommonConfig . Fetch common configuration from etcd
func FetchCommonConfig() {
	log.Printf("Fetching Configuration from %v", Service.RegistryLocation)
	// 'common' sections from etcd
	fetchConfigSection(&Database, DbConfigLocation, kapi)
	fetchConfigSection(&Logger, LogConfigLocation, kapi)
	fetchConfigSection(&ServiceRegistry, RegistryConfigLocation, kapi)
}

// FetchServiceSection Fetch shared configuration for a particular service
func FetchServiceSection(service ServiceSection) {
	switch service {
	case AuthService:
		fetchConfigSection(&Auth, AuthConfigLocation, kapi)
	case UserService:
		fetchConfigSection(&Users, UsersConfigLocation, kapi)
	case WikiService:
		//Do nothing
	default:
		log.Println("Unknown Service config requested")
	}
}

func setConfigVal(str string, field reflect.Value) error {
	t := field.Kind()
	switch {
	case t == reflect.String:
		field.SetString(str)
	case t >= reflect.Int && t <= reflect.Int64:
		if x, err := strconv.ParseInt(str, 10, 64); err != nil {
			return err
		} else {
			field.SetInt(x)
		}
	case t >= reflect.Uint && t <= reflect.Uint64:
		if x, err := strconv.ParseUint(str, 10, 64); err != nil {
			return err
		} else {
			field.SetUint(x)
		}
	case t >= reflect.Float32 && t <= reflect.Float64:
		if x, err := strconv.ParseFloat(str, 64); err != nil {
			return err
		} else {
			field.SetFloat(x)
		}
	case t == reflect.Bool:
		if x, err := strconv.ParseBool(str); err != nil {
			return err
		} else {
			field.SetBool(x)
		}
	default:
		return nil
	}
	return nil
}

//Fetches a single config section
func fetchConfigSection(configStruct interface{}, location string, kapi etcd.KeysAPI) {
	cfg := reflect.ValueOf(configStruct).Elem()
	for i := 0; i < cfg.NumField(); i++ {
		key := cfg.Type().Field(i).Name
		resp, getErr := kapi.Get(context.Background(), location+key, nil)
		if getErr != nil {
			log.Printf("Error getting key %v: %v\n", key, getErr)
			continue
		}
		valErr := setConfigVal(resp.Node.Value, cfg.Field(i))
		if valErr != nil {
			log.Printf("Error setting config field %v: %v\n", key, valErr)
		}
	}
}
