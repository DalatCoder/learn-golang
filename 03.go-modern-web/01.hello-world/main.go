package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintln(rw, "<h1>Home page</h1>")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Number of bytes written to Home page", n)
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintln(rw, "<h1>About page</h1>")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Number of bytes written to About page", n)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting the application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
