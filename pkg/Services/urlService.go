package services

import (
	"math/rand"
	"strings"
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
