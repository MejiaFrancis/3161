package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//create multiplexer
	router := httprouter.New()
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //exclude resource and go to static

	router.HandlerFunc(http.MethodGet, "/create", app.Greeting) //passing in pointer, say where to find handler func
	// callback - above shows passing of the address not the func itself
	router.HandlerFunc(http.MethodGet, "/about", app.About)
	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodPost, "/create", app.MessageCreate)
	//router.HandlerFunc(http.MethodGet,"/poll/reply", app.pollReplyShow)
	//router.HandlerFunc(http.MethodPost, "/poll/reply", app.pollReplySubmit)
	//router.HandlerFunc(http.MethodGet, "/poll/success", app.pollSuccessShow)
	//router.HandlerFunc(http.MethodGet, "/poll/create", app.pollCreateShow)
	//router.HandlerFunc(http.MethodPost, "/poll/create", app.pollCreateSubmit)
	//router.HandlerFunc(http.MethodGet, "/options/create", app.optionsCreateShow)
	//router.HandlerFunc(http.MethodPost, "/options/create", app.optionsCreateSubmit)

	return router
}