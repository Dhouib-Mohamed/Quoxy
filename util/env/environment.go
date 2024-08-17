package env

import (
	"os"
	"strings"
)

const (
	PROD = iota
	TEST
	DEV
)

func GetEnvironment() int {
	value, _ := os.LookupEnv("ENV")

	switch strings.ToLower(value) {
	case "prod":
		return PROD
	case "test":
		return TEST
	case "dev":
		return DEV
	default:
		return DEV
	}
}
