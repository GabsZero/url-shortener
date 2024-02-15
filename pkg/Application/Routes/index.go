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
	router.HandleFunc("/hello", createShortUrlController.Execute).Methods("Get")

	http.ListenAndServe(":8000", router)
}
