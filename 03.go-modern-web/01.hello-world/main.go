package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.tmpl")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting the application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
