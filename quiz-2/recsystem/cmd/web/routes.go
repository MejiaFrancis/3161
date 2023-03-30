// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/about", app.about)
	router.HandlerFunc(http.MethodGet, "/login", app.loginform)
	router.HandlerFunc(http.MethodGet, "/register", app.register)
	router.HandlerFunc(http.MethodGet, "/reservation", app.reserve)
	router.HandlerFunc(http.MethodPost, "/feedback", app.feedback)
	router.HandlerFunc(http.MethodPost, "/user", app.userPortal)
	router.HandlerFunc(http.MethodPost, "/admin", app.adminPortal)

	return router
}
