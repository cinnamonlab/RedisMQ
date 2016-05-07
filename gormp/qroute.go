package gormq

import (
	"gopkg.in/redis.v3"
	"errors"
)

type QRoute struct {
	Functions map[string]QueueFunc
}

type QueueFunc func(input string)

// Singleton Route instance
func NewQRoute() *QRoute {
	routeInstance := &QRoute{}
	return routeInstance
}

func (route *QRoute) AddRoutes(controller QController)  {
	if route.Functions == nil {
		route.Functions = make(map[string]QueueFunc)
	}
	for pattern, function := range controller.Routes() {
		route.Functions[pattern]= function
	}
}

func (route *QRoute) GetPerformTask(msg *redis.Message) (*QueueTask,error) {
	// add to queue and run one by one
	if function, ok := route.Functions[msg.Pattern]; ok {
		return &QueueTask{
			Function:function,
			Param:msg.Payload,
		},nil
	} else {
		return nil,errors.New("pattern is not match")
	}
}


