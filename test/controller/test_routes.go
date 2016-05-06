package controller

import "github.com/cinnamonlab/gormq/gormp"

type TestController struct {
	Functions map[string]gormq.QueueFunc
}

func NewInstance() gormq.QController{
	instance := &TestController{}

	instance.initRoutes()

	return instance
}

func (controller *TestController) Routes() map[string]gormq.QueueFunc  {
	return  controller.Functions
}

func (controller *TestController) initRoutes() {

	controller.Functions = make(map[string]gormq.QueueFunc)

	controller.Functions = map[string]gormq.QueueFunc {
		"cache/*/insert":controller.firstController,
	}
}

