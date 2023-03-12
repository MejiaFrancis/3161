package main

import (
	"net/http"
	//"strconv"

	"github.com/MejiaFrancis/3161/3162/quiz-2/recsystem/helpers"
)

// include --about
// include --home
// create handler for greeting
func (app *application) Greeting(w http.ResponseWriter, r *http.Request) {

	helpers.RenderTemplates(w, "./static/html/poll.page.tmpl")
	//RenderTemplate(w, "home.page.tmpl", nil)
	// w.Write([]byte("Welcome to my page."))
	//question, err := app.question.Get()
	//if err != nil {
	//return
	//}
	//w.Write([]byte(question.Body))
}

// create handler for about
func (app *application) About(w http.ResponseWriter, r *http.Request) {
	// RenderTemplate(w, "about.page.tmpl", nil)
	// day := time.Now().Weekday()
	// w.Write([]byte(fmt.Sprintf("Welcome to my  about page, have a nice %s", day)))
	w.Write([]byte("Hello\n"))
}

// create handler for home
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return

	} // w.Write([]byte("Welcome to my home page."))
	//helpers.RenderTemplates(w, "./static/html/home.page.tmpl")

}

// create handler for home
func (app *application) MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("ALLOW", "POST") //setting header in order to do a 'write'
		//w.WriteHeader(405) //write in header
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		//w.Write([]byte("method not allowed")) //this is writing in the body
		return
	}
	// w.Write([]byte("method created..."))
	// get the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	the_question := r.PostForm.Get("new_question")
	_, err = app.user.Insert(the_question)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
