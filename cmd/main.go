package main

import (
	"fmt"
	"log"

	"github.com/Shepherd-Go/Nlj-Internal-Email-Service/cmd/providers"
	"github.com/Shepherd-Go/Nlj-Internal-Email-Service/config"
	"github.com/Shepherd-Go/Nlj-Internal-Email-Service/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
