package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
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

	// router.HandlerFunc(http.MethodGet, "/login", app.Login)
	// router.HandlerFunc(http.MethodPost, "/login-auth", app.LoginSubmit)
	router.HandlerFunc(http.MethodGet, "/sign-in", app.SignIn)
	router.HandlerFunc(http.MethodPost, "/sign-in-auth", app.SignInSubmit)
	router.HandlerFunc(http.MethodGet, "/scan-qr-code", app.ScanQrCode)
	router.HandlerFunc(http.MethodPost, "/scan-qr-code-check", app.ScanQrCodeSubmit)

	// return router

	router.Handler(http.MethodGet, "/user/signup", dynamicMiddleware.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamicMiddleware.ThenFunc(app.userSignupSubmit))
	router.Handler(http.MethodGet, "/user/login", dynamicMiddleware.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamicMiddleware.ThenFunc(app.userLoginSubmit))
	router.Handler(http.MethodPost, "/user/logout", dynamicMiddleware.ThenFunc(app.userLogoutSubmit))

	// tidy up the middleware chain
	standardMiddleware := alice.New(app.RecoverPanicMiddleware, app.logRequestMiddleware, securityHeadersMiddleware)
	return standardMiddleware.Then(router)
}
