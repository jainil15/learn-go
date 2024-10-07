package config

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Value for env variable %s not found", key)
	}
	return val
}

type config struct {
	DatabaseUrl string
	JwtSecret   string
	Port        string
	Environment string
}

var Env *config

func initConfig() *config {
	if Env == nil {
		lock.Lock()
		defer lock.Unlock()
		if Env == nil {
			fmt.Println("INIT CONFIG")
			Env = &config{
				DatabaseUrl: MustGetEnv("DATABASE_URL"),
				JwtSecret:   MustGetEnv("JWT_SECRET"),
				Port:        MustGetEnv("PORT"),
				Environment: MustGetEnv("ENVIRONMENT"),
			}
		} else {
			fmt.Println("config already initialized")
		}
	} else {
		fmt.Println("config already initialized")
	}
	return Env
}

func GetConfig() *config {
	return initConfig()
}

var Envs = GetConfig()
