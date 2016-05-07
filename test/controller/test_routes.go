package controller

import "github.com/cinnamonlab/RedisMQ"

type TestController struct {
	Functions map[string]redismq.QueueFunc
}

func NewInstance() redismq.QController{
	instance := &TestController{}

	instance.initRoutes()

	return instance
}

func (controller *TestController) Routes() map[string]redismq.QueueFunc  {
	return  controller.Functions
}

func (controller *TestController) initRoutes() {

	controller.Functions = make(map[string]redismq.QueueFunc)

	controller.Functions = map[string]redismq.QueueFunc {
		"cache/*/insert":controller.firstController,
	}
}

