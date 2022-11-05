package configs

import (
	"os"
)

var AppPort = ":" + os.Getenv("APP_PORT")
var GrpcPort = ":" + os.Getenv("GRPC_PORT")
