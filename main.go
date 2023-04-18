package main

import (
	"gopi/config"
	"gopi/server"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server.Init(config.GetAPIConfig())
}
