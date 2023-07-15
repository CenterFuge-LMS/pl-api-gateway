package authservice

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
)

func AddRoutes(cfg *koanf.Koanf, route *gin.Engine) *AuthServiceClient {
	client := AuthServiceClient{
		Client: InitClient(cfg),
	}
	auth := route.Group("v1/auth-service/")
	{
		auth.POST("/register", client.Register)

	}
	return &client

}
