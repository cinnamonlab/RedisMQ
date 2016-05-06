package controller

import "github.com/cinnamonlab/gormq"

type TestController struct {
	Routes map[string]gormq.QueueFunc
}

func NewInstance() *gormq.QController{
	instance := &TestController{}

	instance.initRoutes()

	return instance
}

func (controller *TestController) Routes() map[string]gormq.QueueFunc  {
	return  controller.Routes
}

func (controller *TestController) initRoutes() {

	controller.Routes = map[string]gormq.QueueFunc {
		"cache/insert/user":controller.firstController,
	}
}

