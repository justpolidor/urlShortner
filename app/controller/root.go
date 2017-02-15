package controller

import (
	"strings"
	"net/http"
	"html/template"
	"urlShortner/app/model"
)

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
