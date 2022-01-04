package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/dalatcoder/bookings/pkg/config"
	"github.com/dalatcoder/bookings/pkg/handlers"
	"github.com/dalatcoder/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// Accessible for all code in the `main package`
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// TODO: change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // persist after user close the browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Starting the application on port:", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
