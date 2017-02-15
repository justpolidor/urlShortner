package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"urlShortner/app/controller"
)

func Routes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/{id}", controller.RootEndpoint).Methods("GET")//hash to long
	router.HandleFunc("/create", controller.CreateEndpoint).Methods("POST")// long to short

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/"))))

	return router
}
