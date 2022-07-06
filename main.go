package main

import (
	"fmt"

	"github.com/capitual/valete_ms/config"
	"github.com/capitual/valete_ms/server"
)

func main() {
	fmt.Println("Hello World")
	config.Init()
	svr := server.NewServer()
	svr.Run()
}
