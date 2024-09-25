package main

import (
	"asciiart/serv"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serv.Index)
	http.HandleFunc("/ascii-art", serv.AsciiWeb)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("http//localhost:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

}
