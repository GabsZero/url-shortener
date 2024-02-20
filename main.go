package main

import (
	"log"
	"os"

	routes "github.com/gabszero/url-shortener/pkg/Application/Routes"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
)

func main() {
	// log to custom file
	LOG_FILE := "./myapp_log.txt"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	databaseRepo := repositories.Repository{}
	databaseRepo.StartDabase()

	router := routes.Router{}
	router.StartRoutes()
}
