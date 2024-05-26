package config

import (
	"api-authenticator-proxy/util/log"
	"api-authenticator-proxy/util/network"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

var (
	yamlContent map[string]interface{}
	loadOnce    sync.Once
)

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Info("no .env file found")
	}
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Info("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &yamlContent)
	if err != nil {
		log.Warning("error in reading yaml file")
	}
}
func init() {
	loadOnce.Do(loadConfig)
}

func getConfigVar(obj interface{}, name string) error {
	loadOnce.Do(loadConfig)
	content, ok := yamlContent[name]
	if !ok {
		log.Warning("Config bloc ", name, "is not found")
		return os.ErrNotExist
	}

	contentBytes, err := yaml.Marshal(content)
	if err != nil {
		log.Info("error in reading bloc ", name)
		return err
	}

	err = yaml.Unmarshal(contentBytes, obj)
	if err != nil {
		log.Info("error in reading bloc ", name)
		return err
	}

	return nil
}

func getValidPort(port string) string {
	log.Debug("Provided port : ", port, " is valid : ", network.IsPortValid(port))
	provided := port != ""
	if provided && network.IsPortValid(port) {
		return port
	}
	for i := 8000; i < 8100; i++ {
		if testPort := strconv.Itoa(i); network.IsPortValid(testPort) {
			port = testPort
			break
		}
	}
	if provided {
		log.Warning("Provided port is in use. Using port : ", router.Port)
	}
	return port
}
