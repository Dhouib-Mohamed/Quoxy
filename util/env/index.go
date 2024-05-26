package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

func warning(message string) {
	fmt.Println(" [Warning] :", message)
}

var loadOnce sync.Once

func init() {
	loadOnce.Do(loadEnv)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		warning("no .env file found, using default values")
	}
}

func getEnvVar(key string) string {
	loadOnce.Do(loadEnv)
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	} else {
		warning("Environment variable " + key + " not found, going back to default value")
		return ""
	}
}
