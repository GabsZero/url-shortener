package services

import (
	"errors"
	"log"
	"math/rand"
	"strings"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	repositories "github.com/gabszero/url-shortener/pkg/Infrastructure/Repositories"
)

type UrlService struct {
}

func (us *UrlService) GetShard(firstLetter string) int {

	const firstShardLetters = "abcdefghijklmnopqrstuvwxyzABCDE"
	if strings.Contains(firstShardLetters, firstLetter) {
		return 1
	}

	return 2
}

func (us *UrlService) RandomString(size int) string {
	const alphaNumerical = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

	b := make([]byte, size)
	for i := range b {
		b[i] = alphaNumerical[rand.Intn(len(alphaNumerical))]
	}
	return string(b)
}

func (us *UrlService) CreateRandomShortUrl(url *models.Url, urlSize int) (bool, error) {
	randomString := us.RandomString(urlSize)
	url.Short_url = randomString
	shard := us.GetShard(string(randomString[0]))

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance(shard)

	// checking if short url already exists
	result := db.First(&models.Url{}, "short_url = ?", randomString)

	if result.RowsAffected > 0 {
		for result.RowsAffected > 0 {
			randomString = us.RandomString(urlSize)
			result = db.First(&models.Url{}, "short_url = ?", randomString)
			url.Short_url = randomString
		}
	}

	createResult := db.Create(&url)

	if createResult.Error != nil {
		return false, errors.New("something went wrong while saving the url")
	}

	return true, nil
}

func (us *UrlService) CreateCustomShortUrl(url models.Url) (bool, error) {
	shard := us.GetShard(string(url.Short_url[0]))

	mainRepo := repositories.Repository{}
	db := mainRepo.GetDbInstance(shard)

	// checking if short url already exists
	result := db.First(&models.Url{}, "short_url = ?", url.Short_url)

	if result.RowsAffected > 0 {
		return false, errors.New("custom url provided is already in use")
	}

	createResult := db.Create(&url)

	if createResult.Error != nil {

		log.Println(createResult.Error)
		return false, errors.New("something went wrong while saving the url")
	}

	return true, nil
}
