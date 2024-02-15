package routes

import (
	"net/http"

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
	router.HandleFunc("/{short_url}", redirectController.Execute).Methods("Get")

	http.ListenAndServe(":8000", router)
}
