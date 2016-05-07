# RedisMQ
Message queue broker based on redis, support route also.

<br>

<img src = "https://github.com/cinnamonlab/RedisMQ/blob/master/QueueServer.png"/>

# One queue instance will be create from:
- route pattern which is the same subscribe chanel pattern.
- Controller for each event happen on each chanel (when receive new message....)


# To define controller and route, follow:
- Init route(s) for controller
```go
func (controller *TestController) initRoutes() {

	controller.Functions = make(map[string]gormq.QueueFunc)

	controller.Functions = map[string]gormq.QueueFunc {
		"cache/*/insert":controller.firstController,
	}
}
```
- Create controller function for each route, like this:

```go
func (controller *TestController) firstController(input string)  {
	fmt.Println("Message payload:"+input)
}
```
- Init new connection with all routes defined before start connection.
```go
  route := gormq.NewQRoute()

	testController := controller.NewInstance()

	route.AddRoutes(testController)

	client := gormq.NewConn(route);

	client.Start("127.0.0.1","6379")

```
- When new connection is started, this client is automatic subscribe to all channels with pattern is route paths and start workerPool also.
- Foreach new message received from channels, it will add to workerpool to perform.

```go
func (conn *Conn) subscribes() {

	patterns := make([]string,0)

	for pattern, _ := range conn.Route.Functions {
		patterns = append(patterns,pattern)
	}

	if len(patterns)>0 {
		fmt.Println(patterns)

		pubsub, err := conn.Client.PSubscribe(patterns...)

		if err != nil {
			fmt.Println("subscribe error")
			return 
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
```

