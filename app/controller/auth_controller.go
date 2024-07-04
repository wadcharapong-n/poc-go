package controller

import (
	"github.com/gin-gonic/gin"
	"poc-go/app/service"
)

type AuthController interface {
	Login(c *gin.Context)
	Validate(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (u AuthControllerImpl) Login(c *gin.Context) {
	u.svc.Login(c)
}

func (u AuthControllerImpl) Validate(c *gin.Context) {
	u.svc.Validate(c)
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: authService,
	}
}
