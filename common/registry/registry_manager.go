package registry

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	etcd "github.com/coreos/etcd/client"
	"github.com/laidingqing/amadd9/common/config"
	"golang.org/x/net/context"
)

type nodeCache struct {
	sync.RWMutex
	m map[string][]*etcd.Node
}

var serviceCache = nodeCache{m: make(map[string][]*etcd.Node)}
var pluginCache = nodeCache{m: make(map[string][]*etcd.Node)}

var random = rand.New(rand.NewSource(time.Now().Unix()))

var client *etcd.Client
var kapi etcd.KeysAPI
var EtcdPrefix = "/amadd9"
var UsersLocation = EtcdPrefix + "/services/users/"
var WikisLocation = EtcdPrefix + "/services/wikis/"
var AuthLocation = EtcdPrefix + "/services/auth/"
var ConfigServiceLocation = EtcdPrefix + "/services/config/"

//Config locations
var ConfigPrefix = EtcdPrefix + "/config/"
var DbConfigLocation = ConfigPrefix + "db/"

var ttl time.Duration

func protocolString() string {
	if config.Service.UseSSL {
		return "https://"
	} else {
		return "http://"
	}
}

func GetEtcdKeyAPI() etcd.KeysAPI {
	return kapi
}

func hostUrl() string {
	return protocolString() + config.Service.DomainName +
		":" + config.Service.Port
}

func Init(serviceName, registryLocation string) error {
	log.Println("Initializing registry connection.")
	nodeID := config.Service.NodeId
	ttl = time.Duration(config.ServiceRegistry.EntryTTL) * time.Second
	cfg := etcd.Config{
		Endpoints: []string{config.Service.RegistryLocation},
		Transport: etcd.DefaultTransport,
	}
	client, err := etcd.New(cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	kapi = etcd.NewKeysAPI(client)
	log.Println("Registering " + serviceName + " Service node at " + hostUrl())
	if _, err := kapi.Set(context.Background(), registryLocation+nodeID, hostUrl(),
		&etcd.SetOptions{TTL: ttl}); err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return err
	}
	fetchServiceLists()
	go sendHeartbeat(registryLocation)
	go updateServiceCache()
	return nil
}

func sendHeartbeat(registryLocation string) {
	nodeID := config.Service.NodeId
	for {
		time.Sleep(time.Duration(config.ServiceRegistry.EntryTTL/2) * time.Second)
		if _, err := kapi.Set(context.Background(), registryLocation+nodeID,
			hostUrl(), &etcd.SetOptions{TTL: ttl}); err != nil {
			log.Print("Can't send Heartbeat to registry! - " + err.Error())
		}
	}
}

func updateServiceCache() {
	cri := config.ServiceRegistry.CacheRefreshInterval
	for {
		time.Sleep(time.Duration(cri) * time.Second)
		fetchServiceLists()
	}
}

func getServiceNodes(serviceLocation string) ([]*etcd.Node, error) {
	ctx, _ := context.WithTimeout(context.Background(), 7*time.Second)
	if resp, err := kapi.Get(ctx, serviceLocation, &etcd.GetOptions{Recursive: true}); err != nil {
		return nil, err
	} else {
		return processResponse(resp)
	}

}

// Loads the latest services from Etcd
func fetchServiceLists() {
	// First, fetch the core services
	userNodes, err := getServiceNodes(UsersLocation)
	if err != nil {
		log.Println("Error fetching user services: " + err.Error())
	}
	// wikiNodes, err := getServiceNodes(WikisLocation)
	// if err != nil {
	// 	log.Println("Error fetching wiki services: " + err.Error())
	// }
	authNodes, err := getServiceNodes(AuthLocation)
	if err != nil {
		log.Println("Error fetching auth services: " + err.Error())
	}
	configNodes, err := getServiceNodes(ConfigServiceLocation)
	if err != nil {
		log.Println("Error fetching config services: " + err.Error())
	}
	serviceCache.Lock()
	defer serviceCache.Unlock()
	serviceCache.m["users"] = userNodes
	// serviceCache.m["wikis"] = wikiNodes
	serviceCache.m["auth"] = authNodes
	serviceCache.m["config"] = configNodes

}

//Read nodes from an etcd response
func processResponse(response *etcd.Response) ([]*etcd.Node, error) {
	rootNode := response.Node
	if !rootNode.Dir {
		return nil, errors.New("Not a directory!")
	}
	if len(rootNode.Nodes) == 0 {
		return nil, errors.New("No listed services!")
	}
	return rootNode.Nodes, nil
}

func getEndpointFromNode(node *etcd.Node) string {
	return node.Value
}

//GetServiceLocation Get a service node for use
func GetServiceLocation(serviceName string) (string, error) {
	serviceCache.RLock()
	defer serviceCache.RUnlock()
	if max := len(serviceCache.m[serviceName]); max == 0 {
		return "", errors.New("No " + serviceName + " services listed!")
	} else {
		index := 0
		if max > 1 {
			index = random.Intn(max)
		}
		return getEndpointFromNode(serviceCache.m[serviceName][index]), nil
	}
}
