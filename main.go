package main

import (
	"github.com/capitual/valete_ms/config"
	"github.com/capitual/valete_ms/server"
)

func main() {
	config.Init()
	svr := server.NewServer()
	svr.Run()
}
