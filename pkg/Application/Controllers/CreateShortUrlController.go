package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
	services "github.com/gabszero/url-shortener/pkg/Services"
)

type CreateShortUrlController struct {
	urlService services.UrlService
}

func (controller *CreateShortUrlController) newUrl(long_url string, short_url string) (models.Url, error) {
	if len(short_url) <= 7 {
		return models.Url{}, errors.New("Custom url need to be 8 characters or higher")

	}

	url := models.Url{
		Long_url:    long_url,
		Short_url:   short_url,
		Expire_date: time.Now().AddDate(100, 0, 0), //default value
	}

	return url, nil
}

func (controller *CreateShortUrlController) CustomShortUrl(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	long_url := req.PostFormValue("long_url")
	custom_url := req.PostFormValue("custom_url")

	url, err := controller.newUrl(long_url, custom_url)

	if err != nil {
		response := response(false, err.Error(), map[string]any{
			"custom_url": custom_url,
			"length":     len(custom_url),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	url := models.Url{
		Long_url:    long_url,
		Short_url:   custom_url,
		Expire_date: time.Now().AddDate(100, 0, 0),
	}

	_, err := controller.urlService.CreateCustomShortUrl(url)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		response := response(false, err.Error(), map[string]string{
			"custom_url": custom_url,
		})
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := response(true, "Success!", map[string]string{
		"shor_url": url.Short_url,
	})

	w.Write(response)
}

func (controller *CreateShortUrlController) Execute(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	short_url_length := 7

	long_url := req.PostFormValue("long_url")

	randomString := controller.urlService.RandomString(short_url_length)

	url := models.Url{
		Long_url:    long_url,
		Short_url:   randomString,
		Expire_date: time.Now().AddDate(100, 0, 0),
	}

	shard := controller.urlService.GetShard(string(randomString[0]))

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance(shard)

	// checking if short url already exists
	result := db.First(&models.Url{}, "short_url = ?", randomString)

	if result.RowsAffected > 0 {
		for result.RowsAffected > 0 {
			randomString = controller.urlService.RandomString(short_url_length)
			result = db.First(&models.Url{}, "short_url = ?", randomString)
			url.Short_url = randomString
		}
	}

	createResult := db.Create(&url)

	if createResult.Error != nil {
		log.Println(createResult.Error)

		w.WriteHeader(http.StatusInternalServerError)
		response := response(false, "Something went wrong while saving the url", nil)
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := response(true, "Success!", map[string]string{
		"shor_url": url.Short_url,
	})

	w.Write(response)

}
