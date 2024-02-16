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

	url := models.Url{}
	result := db.First(&url, "short_url = ?", urlParameters["short_url"])

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)

		w.Write(response(false, "Url not found", nil))
		return
	}

	if url.Long_url == "" {
		panic("long url not defined")
	}

	url.Is_used = true
	db.Save(url)

	http.Redirect(w, req, url.Long_url, http.StatusSeeOther)
}
