// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamicMiddleware := alice.New(app.sessionManager.LoadAndSave)
	
	// from here
	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/about", dynamicMiddleware.ThenFunc(app.about))
	router.Handler(http.MethodGet, "/login", dynamicMiddleware.ThenFunc(app.loginform))
	router.Handler(http.MethodPost, "/login", dynamicMiddleware.ThenFunc(app.loginformSubmit))
	router.Handler(http.MethodGet, "/register", dynamicMiddleware.ThenFunc(app.register))
	router.Handler(http.MethodPost, "/register", dynamicMiddleware.ThenFunc(app.registerSubmit))
	router.Handler(http.MethodGet, "/reservation", dynamicMiddleware.ThenFunc(app.reserve))
	router.Handler(http.MethodPost, "/reservation", dynamicMiddleware.ThenFunc(app.reserveFormSubmit))
	router.Handler(http.MethodGet, "/feedback", dynamicMiddleware.ThenFunc(app.feedback))
	router.Handler(http.MethodPost, "/feedback", dynamicMiddleware.ThenFunc(app.feedbackFormSubmit))
	router.Handler(http.MethodGet, "/user", dynamicMiddleware.ThenFunc(app.userPortal))
	router.Handler(http.MethodPost, "/user", dynamicMiddleware.ThenFunc(app.userPortalFormSubmit))

	router.Handler(http.MethodGet, "/admin", dynamicMiddleware.ThenFunc(app.adminPortal))
	router.Handler(http.MethodPost, "/admin", dynamicMiddleware.ThenFunc(app.adminPortalFormSubmit))

	//stop here

	//tidy up the middleware chain
	standardMiddleware := alice.New(
		app.recoverPanicMiddleware,
		app.logRequestMiddleware,
		securityHeadersMiddleware,
	)

	return standardMiddleware.Then(router)
}
