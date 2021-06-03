package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/roadluck/reservaciones/pkg/config"
	"github.com/roadluck/reservaciones/pkg/handlers"
	"github.com/roadluck/reservaciones/pkg/render"
)

var portNumber = ":8080"
var app config.AppConfig

var session *scs.SessionManager

// main es la raiz de la aplicacion
func main() {

	//Cambiar esto si va a produccion
	app.Debug = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true

	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Debug

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("No se puede cargar el cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Servidot en el puerto:", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
