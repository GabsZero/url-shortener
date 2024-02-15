package controllers

import (
	"fmt"
	"net/http"
	"time"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
)

type CreateShortUrlController struct {
}

func (controller *CreateShortUrlController) Execute(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	req.ParseForm()

	long_url := req.PostFormValue("long_url")

	url := models.Url{
		Long_url:    long_url,
		Short_url:   "testing123",
		Expire_date: time.Now(),
	}

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance()

	db.Create(&url)

	fmt.Println(url)

	w.Header().Set("Content-Type", "application/json")
	response := response(true, "Success!", nil)

	w.Write(response)

}
