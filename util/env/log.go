package env

import (
	"strings"
)

const (
	NONE = iota
	ERROR
	INFO
	DEBUG
)

func GetLogLevel() int {
	logLevel := getEnvVar("LOG_LEVEL")
	normalizedLogLevel := strings.ToLower(strings.Trim(logLevel, " "))
	switch normalizedLogLevel {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "error":
		return ERROR
	case "none":
		return NONE
	default:
		if normalizedLogLevel != "" {
			warning("LOG_LEVEL invalid, using default value: info")
		}
		return INFO
	}
}
