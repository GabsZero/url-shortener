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

func (controller *CreateShortUrlController) Execute(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
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
		panic(createResult.Error)
	}

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
