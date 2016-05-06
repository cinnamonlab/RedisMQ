package gormq

type QRoute struct {
	Functions map[string]QueueFunc
}

type QueueFunc func(input interface{})

var routeInstance *QRoute

// Singleton Route instance
func GetQRouteInstance() *QRoute {
	if routeInstance == nil {
		routeInstance = &QRoute{}
	}
	return routeInstance
}

func (route *QRoute) AddRoutes()  {

}


