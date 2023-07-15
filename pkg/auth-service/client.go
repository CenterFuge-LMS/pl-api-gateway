package authservice

import (
	"log"
	"pl-api-gateway/pkg/auth-service/pb"

	"github.com/knadh/koanf"
	"google.golang.org/grpc"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func InitClient(config *koanf.Koanf) pb.AuthServiceClient {

	cc, err := grpc.Dial(config.String("services.authService"), grpc.WithInsecure())
	if err != nil {
		log.Println("error connecting to service", err)
	}
	return pb.NewAuthServiceClient(cc)

}
