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

// P godoc
// @Summary Payment
// @Description Payment to another API
// @Tags payment
// @Accept  json
// @Produce  json
// @Param payment body dto.PaRequest true "Payment Request"
// @Success 200 {object} dto.PaResponse
// @Router /p [post]
func (p PControllerImpl) P(c *gin.Context) {
	p.svc.PA(c)
}
