package main

import (
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
)

func main() {
	databaseRepo := repositories.Repository{}
	databaseRepo.StartDabase()

}
