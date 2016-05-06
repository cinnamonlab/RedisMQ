package controller

import "fmt"

func (controller *TestController) firstController(input string)  {
	fmt.Println("Message payload:"+input)
}
