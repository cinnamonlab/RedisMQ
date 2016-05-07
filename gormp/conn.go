package gormq

import (
	"gopkg.in/redis.v3"
	"fmt"
	"github.com/cinnamonlab/WorkerPool"
)

type Conn struct {
	Client *redis.Client
	Route *QRoute
}

func NewConn(route *QRoute) *Conn  {
	return &Conn{
		Client:nil,
		Route:route,
	}
}

func (conn *Conn) Start(host string, port string) error {
	conn.Client = redis.NewClient(&redis.Options{
		Addr:host+":"+port,
		DB:0,
	})

	result, err := conn.Client.Ping().Result()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("PING Response:"+result);

	conn.subscribes()

	return nil
}

func (conn *Conn) subscribes() {

	patterns := make([]string, 0)

	for pattern, _ := range conn.Route.Functions {
		patterns = append(patterns, pattern)
	}

	if len(patterns) > 0 {
		fmt.Println(patterns)

		pubsub, err := conn.Client.PSubscribe(patterns...)

		if err != nil {
			fmt.Println("subscribe error")
		}
		// start worker pool here
		workerpool.Start(5)

		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				fmt.Println("subscribe error:" + err.Error())
			} else {
				task,err := conn.Route.GetPerformTask(msg)

				if err != nil {
					// not match, ignore this case
				} else  {
					workerpool.AddNewTask(task)
				}
				fmt.Println("receive from:" + msg.Channel + " message:" + msg.Payload + " pattern:" + msg.Pattern)
			}
		}
	} else {
		fmt.Println("chanel patterns is empty!")
	}
}



