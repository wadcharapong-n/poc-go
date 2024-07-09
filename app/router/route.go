package router

import (
	"github.com/gin-gonic/gin"
	"poc-go/config"
	"poc-go/middleware"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/login", init.AuthCtrl.Login)
		auth.GET("/profile", middleware.RequireAuthWithBearerToken, init.AuthCtrl.Validate)

	}
	{
		user := api.Group("/user", middleware.RequireAuthWithBearerTokenRoleAdmin)
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	}
	{
		p := api.Group("/p")
		p.POST("", init.PCtrl.P)
	}

	return router
}
