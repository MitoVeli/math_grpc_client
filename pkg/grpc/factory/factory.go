package factory

import (
	"sync"

	"log"

	"google.golang.org/grpc"
)

var (
	connString  = ""
	connectOnce sync.Once
)

func InitializeRpc(connectionString string, initClient func(c *grpc.ClientConn)) {
	if connectionString == "" {
		panic("gRPC connection string is empty")
	}
	connString = connectionString

	if err := connect(initClient); err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
		return
	}
}

// creates a grpc connection
func connect(f func(c *grpc.ClientConn)) error {
	var err error

	clientConn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	} else {
		f(clientConn)
		log.Printf("Connected to gRPC server: %s", connString)
	}

	return err
}
