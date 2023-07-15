package authservice

import (
	"errors"
	"net/http"
	"pl-api-gateway/pkg/auth-service/pb"

	"github.com/gin-gonic/gin"
)

var RegisterPayload struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int64  `json:"phoneNumber"`
}

func (s *AuthServiceClient) Register(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(RegisterPayload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("bad request data"))
	}
	resp, err := s.Client.Register(ctx, &pb.RegisterRequest{
		FirstName:   RegisterPayload.FirstName,
		LastName:    RegisterPayload.LastName,
		Email:       RegisterPayload.Email,
		Password:    RegisterPayload.Password,
		PhoneNumber: RegisterPayload.PhoneNumber,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, resp)

}

