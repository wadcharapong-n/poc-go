package config

import (
	"poc-go/app/controller"
	"poc-go/app/repository"
	"poc-go/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController
	RoleRepo repository.RoleRepository
	AuthCtrl controller.AuthController
	AuthSvc  service.AuthService
	PCtrl    controller.PController
	PSvc     service.PService
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	authCtrl controller.AuthController,
	authSvc service.AuthService,
	pCtrl controller.PController,
	pSvc service.PService) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		RoleRepo: roleRepo,
		AuthCtrl: authCtrl,
		AuthSvc:  authSvc,
		PCtrl:    pCtrl,
		PSvc:     pSvc,
	}
}
