package config_loader

import (
	"log"
	"reflect"
	"strconv"

	etcd "github.com/coreos/etcd/client"
	"github.com/laidingqing/amadd9/common/config"
	. "github.com/laidingqing/amadd9/common/database"
	"golang.org/x/net/context"
)

var kapi etcd.KeysAPI

//Initialize our etcd connection
func InitRegistry() {
	log.Print("Initializing registry connection.")
	cfg := etcd.Config{
		Endpoints: []string{config.Service.RegistryLocation},
		Transport: etcd.DefaultTransport,
	}
	client, err := etcd.New(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	kapi = etcd.NewKeysAPI(client)
}

//Perform some initialization on the CouchDB database
func InitDatabase() {
	InitDb()
	SetupDb()
}

//Load the configuration into etcd
func SetConfig() {
	log.Println("Setting service registry config")
	setConfigItems(config.ServiceRegistry, config.RegistryConfigLocation)
	log.Println("Setting database config")
	setConfigItems(config.Database, config.DbConfigLocation)
	log.Println("Setting logger config")
	setConfigItems(config.Logger, config.LogConfigLocation)
	log.Println("Setting auth config")
	setConfigItems(config.Auth, config.AuthConfigLocation)
	log.Println("Setting users config")
	setConfigItems(config.Users, config.UsersConfigLocation)
}

//Clear the configuration in etcd
func ClearConfig() {
	kapi.Delete(context.Background(), config.ConfigPrefix,
		&etcd.DeleteOptions{
			Recursive: true,
		})
}

func setConfigItems(configStruct interface{}, configLocation string) {
	cfg := reflect.ValueOf(configStruct)
	for i := 0; i < cfg.NumField(); i++ {
		key := cfg.Type().Field(i).Name
		entry := cfg.Field(i).Interface()
		cfgVal := entryToString(entry)
		log.Printf("Setting Key: %v, Value: %v", key, cfgVal)
		if err := setConfigEntry(configLocation+key, cfgVal); err != nil {
			log.Printf("Error setting config "+key+": %v", err)
		}
	}

}

func setConfigEntry(key string, value string) error {
	_, err := kapi.Set(context.Background(), key, value, nil)
	return err
}

func entryToString(entry interface{}) string {
	field := reflect.ValueOf(entry)
	kind := field.Kind()
	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return strconv.FormatInt(field.Int(), 10)
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return strconv.FormatUint(field.Uint(), 10)
	case kind == reflect.Bool:
		return strconv.FormatBool(field.Bool())
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'E', -1, 64)
	case kind == reflect.String:
		return field.String()
	default:
		return ""

	}
}
