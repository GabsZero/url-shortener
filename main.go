package main

import (
	routes "github.com/gabszero/url-shortener/pkg/Application/Routes"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
)

func main() {
	databaseRepo := repositories.Repository{}
	databaseRepo.StartDabase()

	router := routes.Router{}
	router.StartRoutes()
}
