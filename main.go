package main

import (
	"affirmatios/hospital/app"
	"affirmatios/hospital/internal/aagent"
	"log"

	"github.com/joho/godotenv"
)

// AppName is the application name
var AppName = "Hospital Management App"

// AppVersion is the application version
var AppVersion = "0.1"

func main() {
	// Load the configuration parameters from the `.env` file
	godotenv.Load(".env")
	config := app.GetConfig(AppName, AppVersion)
	// intialize the agent
	aagent.InitAgent(config.GetAriesHost(), config.GetAriesPort())
	if err := app.Run(config); err != nil {
		log.Fatal(err)
	}
}
