package controller

import (
	"github.com/gin-gonic/gin"
	"poc-go/app/service"
)

type PController interface {
	P(c *gin.Context)
	B(c *gin.Context)
	D(c *gin.Context)
	SavePac(c *gin.Context)
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

// B godoc
// @Summary Branch
// @Description getEmployeeListByBranch to another API
// @Tags Branch
// @Produce  json
// @Param payment query dto.BRequest true "request"
// @Success 200 {object} dto.BrResponse
// @Router /b [get]
func (p PControllerImpl) B(c *gin.Context) {
	p.svc.BR(c)
}

// D godoc
// @Summary Delete something
// @Description deleteOwnerClass to another API
// @Tags Branch
// @Produce  json
// @Param payment body dto.DRequest true "request"
// @Success 200 {object} dto.DResponse
// @Router /d [delete]
func (p PControllerImpl) D(c *gin.Context) {
	p.svc.DO(c)
}

// SavePac godoc
// @Summary Package
// @Description Package to another API
// @Tags Package
// @Accept  json
// @Produce  json
// @Param payment body dto.SavePackageRequest true "SavePackageRequest"
// @Success 200 {object} dto.SavePacResponse
// @Router /savePac [post]
func (p PControllerImpl) SavePac(c *gin.Context) {
	p.svc.SavePac(c)
}
