package controllers

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
)

type CreateShortUrlController struct {
}

func (controller *CreateShortUrlController) CustomShortUrl(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	long_url := req.PostFormValue("long_url")
	custom_url := req.PostFormValue("custom_url")
	if len(custom_url) <= 7 {
		response := response(false, "Custom url need to be 8 characters or higher", map[string]any{
			"custom_url": custom_url,
			"lenght":     len(custom_url),
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

	shard := GetShard(string(custom_url[0]))

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance(shard)

	// checking if short url already exists
	result := db.First(&models.Url{}, "short_url = ?", custom_url)

	if result.RowsAffected > 0 {
		// custom url already exists
		response := response(false, "Custom url provided is already in use", map[string]string{
			"custom_url": custom_url,
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	createResult := db.Create(&url)

	if createResult.Error != nil {
		log.Println(createResult.Error)
		panic(createResult.Error)
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

	randomString := RandomString(short_url_length)

	url := models.Url{
		Long_url:    long_url,
		Short_url:   randomString,
		Expire_date: time.Now().AddDate(100, 0, 0),
	}

	shard := GetShard(string(randomString[0]))

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance(shard)

	// checking if short url already exists
	result := db.First(&models.Url{}, "short_url = ?", randomString)

	if result.RowsAffected > 0 {
		for result.RowsAffected > 0 {
			randomString = RandomString(short_url_length)
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

func GetShard(firstLetter string) int {

	const firstShardLetters = "abcdefghijklmnopqrstuvwxyzABCDE"
	if strings.Contains(firstShardLetters, firstLetter) {
		return 1
	}

	return 2
}

func RandomString(size int) string {
	const alphaNumerical = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

	b := make([]byte, size)
	for i := range b {
		b[i] = alphaNumerical[rand.Intn(len(alphaNumerical))]
	}
	return string(b)
}
