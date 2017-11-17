package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	//Handle user flags and set defaults
	port := flag.String("p", "8000", "port to listen on")
	targetDir := flag.String("d", "/tmp", "target directory to serve")
	flag.Parse()

	//setup http handlers
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(*targetDir)))
	http.Handle("/", r)
	fmt.Println("Cargo Bay serving " + *targetDir + " on port " + *port)
	http.ListenAndServe(":"+*port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I suggest we increase the ventilation in the cargo bay before we are asphyxiated.")
}
