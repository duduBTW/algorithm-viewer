package main

import (
	"aba/components"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// PUBLIC SERVER
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public", fileServer))

	// ROUTE HANDLERS
	http.Handle("/", templ.Handler(components.Index(components.AllAlgorithms())))
	for _, page := range components.AllAlgorithmsPages {
		http.Handle(page.Url(), templ.Handler(components.Index(page.Component())))
	}

	// SERVE APP
	http.ListenAndServe(":8080", nil)
}
