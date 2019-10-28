package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var VERSION = "v1.1"

func main() {

	//Handle user flags and set defaults
	port := flag.String("p", "8000", "port to listen on")
	targetDir := flag.String("d", "/tmp", "target directory to serve")
	versionFlag := flag.Bool("v", false, "display version")
	quietFlag := flag.Bool("q", false, "quiet mode")
	flag.Parse()
	if *versionFlag {
		fmt.Println("cargobay " + VERSION)
		return
	}

	//setup http handlers
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(*targetDir)))
	if ! *quietFlag {
		r.Use(loggingMiddleware)
	}
	http.Handle("/", r)
	fmt.Println("Cargo Bay serving " + *targetDir + " on port " + *port)
	http.ListenAndServe(":"+*port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I suggest we increase the ventilation in the cargo bay before we are asphyxiated.")
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //log each request
        log.Println(r.Method, r.URL, r.Host)
        next.ServeHTTP(w, r)
    })
}
