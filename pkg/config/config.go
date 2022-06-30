package config

import (
	"log"

	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

type Config struct {
	JwtSecret        string `json:"jwt_secret"`
	Port             string `json:"port"`
	PostgresHost     string `json:"postgres_host"`
	PostgresUser     string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresDbname   string `json:"postgres_dbname"`
	PostgresPort     string `json:"postgres_port"`
	PostgresSslmode  string `json:"postgres_sslmode"`
	PostgresTimezone string `json:"postgres_timezone"`

	GITHUB_SECRET    string `json:"GITHUB_SECRET"`
	GITHUB_CLIENT_ID string `json:"GITHUB_CLIENT_ID"`
	REDIRECT         string `json:"REDIRECT"`
}

var config Config

func initConfig() map[string]string {
	var err error
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Panic(err)
	}

	return envs
}

func GetConfig() Config {
	onceBody := func() {
		envs := initConfig()

		config = Config{
			JwtSecret:        envs["jwt_secret"],
			Port:             envs["port"],
			PostgresHost:     envs["postgres_host"],
			PostgresUser:     envs["postgres_user"],
			PostgresPassword: envs["postgres_password"],
			PostgresDbname:   envs["postgres_dbname"],
			PostgresPort:     envs["postgres_port"],
			PostgresSslmode:  envs["postgres_sslmode"],
			PostgresTimezone: envs["postgres_timezone"],
			GITHUB_SECRET:    envs["GITHUB_SECRET"],
			GITHUB_CLIENT_ID: envs["GITHUB_CLIENT_ID"],
			REDIRECT:         envs["REDIRECT"],
		}
	}

	once.Do(onceBody)

	return config
}
