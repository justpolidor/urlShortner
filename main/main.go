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
		log.Println("sono qui")
		log.Fatal(err)
	}
	ShortUrl := model.GenerateShortUrl(r.Form.Get("url"), genesisNumber)
	log.Println(ShortUrl)
	tmpl := template.Must(template.ParseFiles("templates/success.html"))
	tmpl.ExecuteTemplate(w, "success.html", string(ShortUrl)) // qui
	//http.Redirect(w,r, "/success", 201)

}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		reqId := strings.Replace(r.URL.Path, "/", "", -1)
		if res, err := model.GetLongUrl(reqId); err != nil {
			log.Fatal(err)
		} else {
			log.Println("redirect")
			http.Get(res)
			return
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
