package controller

import (
	"github.com/gin-gonic/gin"
	"poc-go/app/service"
)

type PController interface {
	P(c *gin.Context)
}

type PControllerImpl struct {
	svc service.PService
}

func PControllerInit(pService service.PService) *PControllerImpl {
	return &PControllerImpl{
		svc: pService,
	}
}

func (p PControllerImpl) P(c *gin.Context) {
	p.svc.PA(c)
}
