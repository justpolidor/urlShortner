package controller

import (
	"log"
	"html/template"
	"net/http"
	"urlShortner/app/model"
)

func CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	var myUrl model.Url
	r.ParseForm()
	if err := model.MatchUrl(r.Form.Get("url")); err != nil {
		tmpl := template.Must(template.ParseFiles("templates/failed.html"))
		tmpl.ExecuteTemplate(w, "failed.html", err.Error())
	} else {
		ShortUrl := myUrl.GenerateShortUrl(r.Form.Get("url"), 6)
		log.Println(ShortUrl)
		tmpl := template.Must(template.ParseFiles("templates/success.html"))
		tmpl.ExecuteTemplate(w, "success.html", string(ShortUrl))
	}

}
