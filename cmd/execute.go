package cmd

import (
	"ctracker/apicall"
	"fmt"
)

func Execute() {
	fmt.Println("Hello, World!")
	// New Service
	// service.NewService()
	// // redis test
	// redisdb.TestRedis("key", "Value")
	apicall.ServeApi()
}
