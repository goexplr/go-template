package main

import (
	"echo-template/app/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	utils.InitLogger()
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Error(err)
	}
	if os.Getenv("ENV") == "development" {
		utils.DebugLogging()
	}

	utils.Logger.Info("Hello, World!")
}
