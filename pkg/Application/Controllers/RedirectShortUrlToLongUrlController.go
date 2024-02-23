package controllers

import (
	"log"
	"net/http"
	"time"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
	services "github.com/gabszero/url-shortener/pkg/Services"
	"github.com/gorilla/mux"
)

type RedirectShortUrlToLongUrlController struct {
	urlService services.UrlService
}

func (controller *RedirectShortUrlToLongUrlController) Execute(w http.ResponseWriter, req *http.Request) {
	urlParameters := mux.Vars(req)

	mainRepo := repositories.Repository{}
	shard := controller.urlService.GetShard(string(urlParameters["short_url"][0]))
	db := mainRepo.GetDbInstance(shard)
	redisInstance := mainRepo.GetRedisInstance()

	url := models.Url{}
	long_url, noValueFoundError := redisInstance.Get(req.Context(), urlParameters["short_url"]).Result()

	if noValueFoundError != nil {
		long_url = ""
	}

	if long_url != "" {
		log.Println("Found url in cache")
		http.Redirect(w, req, long_url, http.StatusSeeOther)
		return
	}

	result := db.First(&url, "short_url = ? and expire_date > ?", urlParameters["short_url"], time.Now())

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)

		w.Write(response(false, "Url not found", nil))
		return
	}

	err := redisInstance.Set(req.Context(), url.Short_url, url.Long_url, 0).Err()
	if err != nil {
		panic(err)
	}

	url.Is_used = true
	db.Save(url)

	http.Redirect(w, req, url.Long_url, http.StatusSeeOther)
}
