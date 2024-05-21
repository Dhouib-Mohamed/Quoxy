package env

import (
	"api-authenticator-proxy/util/log"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
)

var yamlFile []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Info("no .env file found")
	}

	yamlFile, err = ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Info(fmt.Sprintf("yamlFile.Get err #%v ", err))
	}
}

func getEnvVar(key string) (string, error) {
	value, exists := os.LookupEnv(key)

	if exists {
		return value, nil
	} else {
		return "", fmt.Errorf("key %s not found", key)
	}
}

func getConfigVar(obj map[string]interface{}) error {
	return yaml.Unmarshal(yamlFile, obj)
}

func GetPort() int {
	port, err := getEnvVar("SMOKE_TEST_PORT")
	if err != nil {
		return 3000
	}
	value, err := strconv.Atoi(port)
	if err != nil {
		return 3000
	}
	return value
}
