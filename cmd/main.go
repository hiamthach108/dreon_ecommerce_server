package main

import (
	"dreon_ecommerce_server/pkg/adapters"
	"dreon_ecommerce_server/pkg/infrastrutures/server"
)

func init() {
	// This will be executed first
	adapters.IocConfigs()
	adapters.IoCMapper()
	adapters.IoCCache()
	adapters.IoCLogger()
	adapters.IoCCrypto()
	adapters.IoCDatabase()
}

func main() {
	// This will be executed second
	server.StartEchoServer()
}
