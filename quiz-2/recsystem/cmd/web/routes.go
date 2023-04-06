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
	router.HandlerFunc(http.MethodPost, "/login", app.loginform)// function needs to be creates to authenticate user identity
	router.HandlerFunc(http.MethodGet, "/register", app.register)
	router.HandlerFunc(http.MethodPost, "/register", app.register) // function needs to be created to save registration data 
	router.HandlerFunc(http.MethodGet, "/reservation", app.reserve)
	router.HandlerFunc(http.MethodPost, "/reservation", app.reserve) // function needed to be created to save reservation (create handler)
	router.HandlerFunc(http.MethodGet, "/feedback", app.feedback)
	router.HandlerFunc(http.MethodPost, "/feedback", app.feedback)// function need to be created to save the feedback given 

	router.HandlerFunc(http.MethodPost, "/user", app.userPortal)
	router.HandlerFunc(http.MethodPost, "/admin", app.adminPortal)

	return router
}
