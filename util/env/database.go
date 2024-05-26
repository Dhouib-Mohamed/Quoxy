package env

func GetDatabasePath() string {
	path := getEnvVar("DATABASE_PATH")
	if path == "" {
		return "db.sqlite"
	}
	return path
}

func GetDatabaseInitFile() string {
	initFile := getEnvVar("DATABASE_INIT_FILE")
	if initFile == "" {
		return "init.sql"
	}
	return initFile
}
