package main

import (
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/app"
	"github.com/sugandhasaxena1911/GoSamples/BankingApp/logger"
)

func main() {
	logger.Info("Starting the application")

	app.Start()
}
