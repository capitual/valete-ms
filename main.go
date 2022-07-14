package main

import (
	"github.com/capitual/valete_ms/config"
	infra "github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/server"
)

func main() {
	config.Init()
	infra.StartDatabase()
	svr := server.NewServer()
	svr.Run()
}
