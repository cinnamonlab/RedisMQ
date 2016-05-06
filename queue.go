package main

import "github.com/cinnamonlab/gormq/gormp"

func main() {
	gormq.NewConn().Start("127.0.0.1","6379")
}
