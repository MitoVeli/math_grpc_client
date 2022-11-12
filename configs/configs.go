package configs

import (
	"os"
)

var AppPort = ":8008"
var GrpcPort = ":50052"
var AppEnv = os.Getenv("APP_ENV")

func SetEnv() string {

	if os.Getenv("APP_ENV") == "test" {
		return "math_server"
	} else {
		return "localhost"
	}
}
