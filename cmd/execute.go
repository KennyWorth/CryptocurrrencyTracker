package cmd

import (
	// "ctracker/service"
	cs "ctracker/grpc"
	"fmt"
)

func Execute() {
	fmt.Println("Hello, World!")
	// New Service
	cs.NewGrpcServer()
	// // redis test
	// redisdb.TestRedis("key", "Value")
	// apicall.ServeApi()
}
