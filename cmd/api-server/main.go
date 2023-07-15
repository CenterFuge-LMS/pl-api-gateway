package main

import (
	"pl-api-gateway/configs"
	auth "pl-api-gateway/pkg/auth-service"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.NewConfig()
	engine := gin.Default()
	auth.AddRoutes(config, engine)
	if err := engine.Run(config.String("dev.port")); err != nil {

		panic("can't start the server")
	}

}
