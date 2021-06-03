package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})

}

//NpSurf adds CSRF protection.
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.Debug,
		SameSite: http.SameSiteLaxMode,
	})

	return crsfHandler

}

//SessionLoad carga y guarda la session con cada request.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
