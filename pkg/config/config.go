package config

import (
	"log"

	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

type Config struct {
	JWT_SECRET string `json:"jwt_secret"`
	PORT       string `json:"port"`
	MYSQL      string `json:"mysql"`
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
			JWT_SECRET: envs["JWT_SECRET"],
			PORT:       envs["PORT"],
			MYSQL:      envs["MYSQL"],
		}
	}

	once.Do(onceBody)

	return config
}
