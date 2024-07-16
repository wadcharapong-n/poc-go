package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"poc-go/config"
	_ "poc-go/docs"
	"poc-go/middleware"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	{
		p := api.Group("/b")
		p.GET("", init.PCtrl.B)
	}

	return router
}
