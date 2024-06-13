package env

import "strings"

func GetDatabasePath() string {
	path := getEnvVar("DATABASE_PATH")
	if path == "" {
		return "db.sqlite"
	}
	return path
}

func GetDatabaseInitFile() string {
	initFile := getEnvVar("DATABASE_INIT_FOLDER")
	if initFile == "" {
		return "scripts/sql/"
	}
	if !strings.HasSuffix(initFile, "/") {
		initFile += "/"
	}
	return initFile
}
