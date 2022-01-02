package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting the application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
