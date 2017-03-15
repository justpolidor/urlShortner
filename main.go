package main

import (
	"net/http"
	"log"
	"urlShortner/app/route"

	"os"
)


func main() {
	log.Println("Started...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), route.Routes()))
	log.Fatal(http.ListenAndServe(":80", route.Routes()))

}
