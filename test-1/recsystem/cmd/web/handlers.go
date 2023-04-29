// handlers.go
package main

import (
	"net/http"
	//"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl", nil)

}
func (app *application) loginform(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.page.tmpl", nil)

}

//loginformSubmit

func (app *application) loginformSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.page.tmpl", nil)

}

//logout

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.page.tmpl", nil)

}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "register.page.tmpl", nil)

}

//registerSubmit

func (app *application) registerSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "register.page.tmpl", nil)

}

func (app *application) reserve(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "reservation.page.tmpl", nil)

}

//reserveFormSubmit

func (app *application) reserveFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "reservation.page.tmpl", nil)

}

func (app *application) userPortal(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "userPortal.page.tmpl", nil)

}

// userPortalFormSubmit
func (app *application) userPortalFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "userPortal.page.tmpl", nil)

}

func (app *application) adminPortal(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "adminPortal.page.tmpl", nil)

}
func (app *application) adminPortalFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "adminPortal.page.tmpl", nil)

}
func (app *application) feedback(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "feedback.page.tmpl", nil)

}

// feedbackFormSubmit
func (app *application) feedbackFormSubmit(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "feedback.page.tmpl", nil)

}

// func (app *application) pollReplyShow(w http.ResponseWriter, r *http.Request) {
// 	reservation, err := app.questions.Get()
// 	if err != nil {
// 		return
// 	}
// 	data := &templateData{
// 		Reservation: reservation,
// 	}
// 	RenderTemplate(w, "poll.page.tmpl", data)
// }

// func (app *application) submitlogin(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "bad request", http.StatusBadRequest)
// 		return
// 	}
// 	response := r.PostForm.Get("emotion")
// 	question_id := r.PostForm.Get("id")
// 	identifier, err := strconv.ParseInt(question_id, 10, 64)
// 	if err != nil {
// 		return
// 	}
// 	_, err = app.responses.Insert(response, identifier)
// 	if err != nil {
// 		http.Error(w,
// 			http.StatusText(http.StatusInternalServerError),
// 			http.StatusInternalServerError)
// 		return
// 	}
// }

// This code stores the different options in a variabe then stores using the Insert function... for radio button.
// func (app *application) optionsCreateSubmit(w http.ResponseWriter, r *http.Request) {
// 	// get the four options
// 	r.ParseForm()
// 	option_1 := r.PostForm.Get("option_1")
// 	option_2 := r.PostForm.Get("option_2")

// 	// save the options
// 	_, err := app.reservations.Insert(option_1, option_2,)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 	}
// }
