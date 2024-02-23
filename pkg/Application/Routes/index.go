package routes

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/gabszero/url-shortener/pkg/Application/Controllers"
	"github.com/gorilla/mux"
)

type Router struct {
}

func (r *Router) StartRoutes() {
	router := mux.NewRouter()

	createShortUrlController := controllers.CreateShortUrlController{}
	redirectController := controllers.RedirectShortUrlToLongUrlController{}
	router.HandleFunc("/shorten-url", createShortUrlController.Execute).Methods("Post")
	router.HandleFunc("/custom-short-url", createShortUrlController.CustomShortUrl).Methods("Post")
	router.HandleFunc("/{short_url}", redirectController.Execute).Methods("Get")
	port := os.Getenv("URL_SHORTENER_HOST_PORT")

	fmt.Println("Listening request on port " + port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
