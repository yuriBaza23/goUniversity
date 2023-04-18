package main

import (
	"gopi/cmd/server"
	"gopi/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server.HttpInit(config.GetAPIConfig())
}
