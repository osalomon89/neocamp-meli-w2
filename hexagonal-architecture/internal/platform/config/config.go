package config

import (
	"os"

	"github.com/joho/godotenv"
)

var configs struct {
	mysqlUser     string
	mysqlPassword string
	mysqlHost     string
	mysqlDB       string
}

func LoadConfigs() error {
	// TODO: validar para variables en produccion.
	return loadENVConfigs()
}

func loadENVConfigs() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	configs.mysqlUser = os.Getenv("MYSQL_USER")
	configs.mysqlPassword = os.Getenv("MYSQL_PASS")
	configs.mysqlHost = os.Getenv("MYSQL_HOST")
	configs.mysqlDB = os.Getenv("MYSQL_DB")

	return nil
}

func GetMySQLUser() string {
	return configs.mysqlUser
}

func GetMySQLPassword() string {
	return configs.mysqlPassword
}

func GetMySQLHost() string {
	return configs.mysqlHost
}

func GetMySQLDatabase() string {
	return configs.mysqlDB
}
