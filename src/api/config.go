package api

import (
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	RedisHost string
	RedisPort string
	RedisUser string
	RedisPassword string
}

func LoadConfig() Config {
	return Config{
		RedisHost:     GetEnvOrPanic("REDIS_HOST"),
		RedisPort:     GetEnvOrPanic("REDIS_PORT"),
		RedisUser:     GetEnvOrPanic("REDIS_USER"),
		RedisPassword: GetEnvOrPanic("REDIS_PASSWORD"),
	}
}

func GetEnvOrPanic(key string) string {
	val, present := os.LookupEnv(key)
	if !present {
		log.Panic().Msgf("ENV var %s not set", key)
	}
	return val
}
