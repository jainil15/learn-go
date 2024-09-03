package config

import (
	"log"
	"os"
)

func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Value for env variable %s not found", key)
	}
	return val
}

type Config struct {
	DatabaseUrl string
}

func initConfig() Config {
	return Config{
		DatabaseUrl: MustGetEnv("DATABASE_URL"),
	}
}

var Envs = initConfig()
