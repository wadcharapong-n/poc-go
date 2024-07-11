package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"poc-go/config"
	_ "poc-go/docs"
	"poc-go/middleware"
)

// @title Example API
// @version 1.0
// @description This is a sample server for a poc bff
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
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

	return router
}
