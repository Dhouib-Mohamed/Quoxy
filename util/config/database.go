package config

import "api-authenticator-proxy/util/log"

type DatabaseEnv struct {
	External bool   `yaml:"external"`
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

var database DatabaseEnv

func init() {
	database = DatabaseEnv{}
	getConfigVar(&database, "database")
}

func IsDatabaseExternal() bool {
	if !database.External {
		log.Warning("Database is not external. Using By default the local sqlite database.")
	}
	return database.External
}

func GetDatabaseDriver() string {
	if database.Driver == "" || (database.Driver != "sqlite" && database.Driver != "postgres" && database.Driver != "mysql" && database.Driver != "oracle" && database.Driver != "mssql") {
		log.Warning("Database driver not set. Please set the driver in config.yaml. Using sqlite by default.")
		return "sqlite"
	}
	return database.Driver

}

func GetDatabaseHost() string {
	if database.Host == "" {
		log.Warning("Database host not set. Please set the host in config.yaml")
		return ""
	}
	return database.Host
}

func GetDatabasePort() string {
	if database.Port == "" {
		log.Warning("Database port not set. Please set the port in config.yaml")
		return "5432"
	}
	return database.Port
}

func GetDatabaseUser() string {
	if database.User == "" {
		log.Warning("Database user not set. Please set the user in config.yaml")
		return "postgres"
	}
	return database.User
}

func GetDatabasePassword() string {
	if database.Password == "" {
		log.Warning("Database password not set. Please set the password in config.yaml")
		return "password"
	}
	return database.Password
}

func GetDatabaseName() string {
	if database.Name == "" {
		log.Warning("Database name not set. Please set the name in config.yaml")
		return "mydb"
	}
	return database.Name
}
