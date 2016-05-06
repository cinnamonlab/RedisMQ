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

func (route *QRoute) PerformMessage(msg *redis.Message) error {
	// add to queue and run one by one
	if function, ok := route.Functions[msg.Pattern]; ok {
		function(msg.Payload)
		return nil
	} else {
		return errors.New("pattern is not match")
	}
}


