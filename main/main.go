package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"strings"
	"urlShortner/model"
	"html/template"
)
const genesisNumber = 6

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if err := model.MatchUrl(r.Form.Get("url")); err != nil {
		tmpl := template.Must(template.ParseFiles("templates/failed.html"))
		tmpl.ExecuteTemplate(w, "failed.html", err.Error())
	} else {
		ShortUrl := model.GenerateShortUrl(r.Form.Get("url"), genesisNumber)
		log.Println(ShortUrl)
		tmpl := template.Must(template.ParseFiles("templates/success.html"))
		tmpl.ExecuteTemplate(w, "success.html", string(ShortUrl))
	}

}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	var longUrl string
	var err error
	if r.URL.Path != "/" && r.URL.Path != "/create" {
		reqId := strings.Replace(r.URL.Path, "/", "", -1)
		if longUrl, err = model.GetLongUrl(reqId); err != nil {
			tmpl := template.Must(template.ParseFiles("templates/failed.html"))
			tmpl.ExecuteTemplate(w, "failed.html", err.Error())
		} else {
			http.Redirect(w,r,longUrl,http.StatusMovedPermanently)
		}
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/{id}", RootEndpoint).Methods("GET")//hash to long
	router.HandleFunc("/create", CreateEndpoint).Methods("POST")// long to short

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/"))))
	router.PathPrefix("/success").Handler(http.StripPrefix("/", http.FileServer(http.Dir("templates/success.html"))))

	log.Println("Started...")
	log.Fatal(http.ListenAndServe(":80", router))

}
