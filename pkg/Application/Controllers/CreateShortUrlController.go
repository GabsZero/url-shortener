package controllers

import (
	"math/rand"
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

	randomString := RandomString(7)

	url := models.Url{
		Long_url:    long_url,
		Short_url:   randomString,
		Expire_date: time.Now().AddDate(100, 0, 0),
	}

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance()

	// checking if short url already exists
	result := db.First(&models.Url{}, "short_url = ?", randomString)

	if result.RowsAffected > 0 {
		for result.RowsAffected > 0 {
			randomString = RandomString(7)
			result = db.First(&models.Url{}, "short_url = ?", randomString)
			url.Short_url = randomString
		}
	}

	db.Create(&url)

	w.Header().Set("Content-Type", "application/json")
	response := response(true, "Success!", map[string]string{
		"shor_url": url.Short_url,
	})

	w.Write(response)

}

func RandomString(size int) string {
	const alphaNumerical = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

	b := make([]byte, size)
	for i := range b {
		b[i] = alphaNumerical[rand.Intn(len(alphaNumerical))]
	}
	return string(b)
}
