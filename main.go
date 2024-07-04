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

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	_ = app.Run(":" + port)
}
