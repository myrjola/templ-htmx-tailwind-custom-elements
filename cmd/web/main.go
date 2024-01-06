package main

import (
	"github.com/a-h/templ"
	"github.com/myrjola/templ-htmx-tailwind-custom-elements/components"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", templ.Handler(components.Base()))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	addr := "127.0.0.1:8080"
	log.Printf("starting server on %s", addr)
	err := http.ListenAndServe(addr, mux)
	log.Fatal("ListenAndServe: ", err)
}
