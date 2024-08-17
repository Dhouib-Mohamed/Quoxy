package config

import (
	"api-authenticator-proxy/util/env"
	"api-authenticator-proxy/util/log"
	"api-authenticator-proxy/util/network"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"sync"
)

var (
	yamlContent map[string]interface{}
	loadOnce    sync.Once
)

func loadConfig() {
	environment := env.GetEnvironment()
	var yamlFile []byte
	var err error
	if environment == env.TEST {
		yamlFile, err = os.ReadFile("../../config.test.yaml")
		if err != nil {
			log.Info("Test yamlFile.Get err #", err)
		}
	} else {
		yamlFile, err = os.ReadFile("config.yaml")
		if err != nil {
			log.Info("yamlFile.Get err #", err)
		}
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
