package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintln(rw, "<h1>Hello World</h1>")

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Number of bytes written", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
