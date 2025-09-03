package handlers

import (
	"github.com/dyeghocunha/udemy"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	main.renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	main.renderTemplate(w, "about.page.tmpl")
}
