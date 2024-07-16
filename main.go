package main

import (
	"github.com/joho/godotenv"
	"os"
	"poc-go/app/router"
	"poc-go/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	config.InitLog()
}

// Init @title Swagger Example API
// @description This is a sample server
// @version 1.0
// @BasePath /api
// @host localhost:8080
// @schemes http
func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	_ = app.Run(":" + port)
}
