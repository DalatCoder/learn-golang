package main

import (
	"net/http"
)

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.tmpl")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.tmpl")
}
