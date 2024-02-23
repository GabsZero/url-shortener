package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	services "github.com/gabszero/url-shortener/pkg/Services"
)

type CreateShortUrlController struct {
	urlService services.UrlService
}

func (controller *CreateShortUrlController) newUrl(long_url string, short_url string, isCustom bool) (models.Url, error) {
	short_url_length := 7
	if isCustom && len(short_url) <= short_url_length {
		return models.Url{}, errors.New("custom url need to be 8 characters or higher")
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

	url, newUrlError := controller.newUrl(long_url, custom_url, true)

	if newUrlError != nil {
		response := response(false, newUrlError.Error(), map[string]any{
			"custom_url": custom_url,
			"length":     len(custom_url),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	_, customUrlError := controller.urlService.CreateCustomShortUrl(url)

	if customUrlError != nil {
		log.Println(customUrlError)
		w.WriteHeader(http.StatusBadRequest)
		response := response(false, customUrlError.Error(), map[string]string{
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

	long_url := req.PostFormValue("long_url")

	url, newUrlError := controller.newUrl(long_url, "", false)
	if newUrlError != nil {
		response := response(false, newUrlError.Error(), map[string]any{
			"long_url": long_url,
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	_, err := controller.urlService.CreateRandomShortUrl(&url, 7) // need to find a better way to do this

	if err != nil {
		log.Println(err)

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
