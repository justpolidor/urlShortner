package main

import (
	"net/http"
	"log"
	"runtime"
	"urlShortner/app/route"
	"os"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	log.Println("Started...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), route.Routes()))
	//log.Fatal(http.ListenAndServe(":80", route.Routes()))

}
