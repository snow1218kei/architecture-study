package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

const (
	DBKey = "conn"
)

type ENV struct {
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     int    `envconfig:"DB_PORT" default:"5432"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
	GOEnv      string `envconfig:"GO_ENV" default:"local"`
	RDBMaxIdle int    `envconfig:"RDB_MAX_IDLE" default:"10"`
	RDBMaxConn int    `envconfig:"RDB_MAX_CONN" default:"30"`
}

var Env ENV

func init() {
	env := os.Getenv("GO_ENV")

	if err := envconfig.Process(env, &Env); err != nil {
		panic(err)
	}
}

func IsLocal() bool {
	return Env.GOEnv == "local"
}
