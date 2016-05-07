package main

import (
	"github.com/cinnamonlab/RedisMQ"
	"github.com/cinnamonlab/RedisMQ/test/controller"
)

func main() {

	// init route and pattern channels for this client
	route := redismq.NewQRoute()

	testController := controller.NewInstance()

	route.AddRoutes(testController)

	client := redismq.NewConn(route);

	client.Start("127.0.0.1","6379")
}
