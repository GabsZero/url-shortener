package controllers

import (
	"net/http"
	"time"

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
	result := db.First(&url, "short_url = ? and expire_date > ?", urlParameters["short_url"], time.Now())

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)

		w.Write(response(false, "Url not found", nil))
		return
	}

	url.Is_used = true
	db.Save(url)

	http.Redirect(w, req, url.Long_url, http.StatusSeeOther)
}
