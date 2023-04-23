package main

import (
	"log"
	"net/http"

	"github.com/MejiaFrancis/3161/3162/test-1/recsystem/helpers"
	//"strconv"
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
	email := r.PostForm.Get("email")
	first_name := r.PostForm.Get("first_name")
	age := r.PostForm.Get("age")
	last_name := r.PostForm.Get("last_name")
	address := r.PostForm.Get("address")
	phone_number := r.PostForm.Get("phone_number")
	roles_id := r.PostForm.Get("roles_id")
	password := r.PostForm.Get("password")
	log.Printf("%s %s %s %s %s %s %s %d %t\n", email, first_name, last_name, age, address, phone_number, roles_id, password)
	userid, err := app.user.Insert(email, first_name, last_name, age, address, phone_number, roles_id, password)
	log.Printf("%s %s %s %s %s %s %s %s %d\n", email, first_name, last_name, age, address, phone_number, roles_id, password, userid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
