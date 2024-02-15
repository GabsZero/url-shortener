package controllers

import (
	"net/http"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
	"github.com/gorilla/mux"
)

type RedirectShortUrlToLongUrlController struct {
}

func (controller *RedirectShortUrlToLongUrlController) Execute(w http.ResponseWriter, req *http.Request) {
	urlParameters := mux.Vars(req)

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance()

	url := models.Url{
		Short_url: urlParameters["short_url"],
	}
	db.Find(&url)

	if url.Long_url == "" {
		panic("long url not defined")
	}

	http.Redirect(w, req, url.Long_url, http.StatusSeeOther)
}
