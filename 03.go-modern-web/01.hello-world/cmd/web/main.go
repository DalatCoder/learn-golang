package main

import (
	"fmt"
	"github.com/dalatcoder/go-beginner/pkg/config"
	"github.com/dalatcoder/go-beginner/pkg/handlers"
	"github.com/dalatcoder/go-beginner/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting the application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
