package main

import (
	"github.com/meanii/imgax/config"
	"github.com/meanii/imgax/router"
	"github.com/meanii/imgax/server"
)

func main() {
	config.InitEnv()
	s := server.InitServer()
	r := router.Router{}
	r.Init()

	s.StartServer()
}
