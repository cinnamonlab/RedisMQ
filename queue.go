package main

import (
	"github.com/cinnamonlab/gormq/gormp"
	"github.com/cinnamonlab/gormq/test/controller"
)

func main() {

	// init route and pattern channels for this client
	route := gormq.NewQRoute()

	testController := controller.NewInstance()

	route.AddRoutes(testController)

	client := gormq.NewConn(route);

	client.Start("127.0.0.1","6379")
}
