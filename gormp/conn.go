package gormq

import (
	"gopkg.in/redis.v3"
	"fmt"
)

type Conn struct {
	Client *redis.Client
	Route QRoute
}

func NewConn() *Conn  {
	return &Conn{}
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
	pubsub, err := conn.Client.PSubscribe("/cache/*/insert")

	if err != nil {
		fmt.Println("subscribe error")
	}
	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			fmt.Println("subscribe error:" + err.Error())
		}

		fmt.Println("receive from:"+msg.Channel+" message:"+msg.Payload+" pattern:"+msg.Pattern)
	}
}



