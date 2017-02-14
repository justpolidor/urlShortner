package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"strings"
	"urlShortner/model"
)


func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if err := model.MatchUrl(r.Form.Get("url")); err != nil {
		log.Fatal(err)
	}
	
}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	reqUrl := strings.Replace(r.URL.Path, "/", "", -1)
	fmt.Fprint(w, reqUrl)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/{id}", RootEndpoint).Methods("GET")//hash to long
	router.HandleFunc("/create", CreateEndpoint).Methods("POST")// long to short
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("views/"))))
	log.Println("Started...")
	log.Fatal(http.ListenAndServe(":80", router))

}
