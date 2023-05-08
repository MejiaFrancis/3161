package main

import (
	"errors"
	"net/http"

	"gibhub.com/MejiaFrancis/3161/3162/test-1/recsystem/internal/models"
	"github.com/justinas/nosurf"
)

// handler for manage equipment
func (app *application) ManageEquipment(w http.ResponseWriter, r *http.Request) {

	data := &templateData{
		CSRFToken: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "equipment-management.page.tmpl", data)

}

func (app *application) chooseRoleShow(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	data := &templateData{
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "role.page.tmpl", data)
}

func (app *application) chooseRoleSubmit(w http.ResponseWriter, r *http.Request) {
	// get the four options
	r.ParseForm()
	RoleStutent := r.PostForm.Get("RoleStudent")
	RoleAdministrator := r.PostForm.Get("RoleAdministrator")
	RoleTeacher := r.PostForm.Get("RoleTeacher")

	// save the roles
	_, err := app.roles.Insert(RoleStutent, RoleAdministrator, RoleTeacher)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

// include --about
// include --home
// create handler for greeting
func (app *application) Greeting(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/html/viewequipment.html")
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
		// }
		// email := r.PostForm.Get("email")
		// first_name := r.PostForm.Get("first_name")
		// age := r.PostForm.Get("age")
		// last_name := r.PostForm.Get("last_name")
		// address := r.PostForm.Get("address")
		// phone_number := r.PostForm.Get("phone_number")
		// roles_id := r.PostForm.Get("roles_id")
		// password := r.PostForm.Get("password")
		// log.Printf("%s %s %s %s %s %s %s %d %t\n", email, first_name, last_name, age, address, phone_number, roles_id, password)
		// userid, err := app.user.Insert(email, first_name, last_name, age, address, phone_number, roles_id, password)
		// log.Printf("%s %s %s %s %s %s %s %s %d\n", email, first_name, last_name, age, address, phone_number, roles_id, password, userid)
		// if err != nil {
		// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		// 	return
	}
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	// remove the entry from the session manager
	flash := app.sessionManager.PopString(r.Context(), "flash")
	data := &templateData{
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "signup.page.tmpl", data)
}

func (app *application) userSignupSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	// write the data to the dable
	err := app.users.Insert(name, email, password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			RenderTemplate(w, "signup.page.tmpl", nil)
		}
		return
	}
	app.sessionManager.Put(r.Context(), "flash", "Signup was successful")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

// create handler for login
func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {

	flash := app.sessionManager.PopString(r.Context(), "flash")
	data := &templateData{
		Flash:     flash,
		CSRFToken: nosurf.Token(r),
	}
	RenderTemplate(w, "login.html", data)

}

// create handler for LoginSubmit
func (app *application) userLoginSubmit(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	// write the data to the dable
	id, err := app.users.Authenticate(email, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			RenderTemplate(w, "login.html", nil)
		}
		return
	}
	// add the user to the session cookie
	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	// add and authenticate entry
	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/user/admin/manage-equipment", http.StatusSeeOther)

}

func (app *application) userLogoutSubmit(w http.ResponseWriter, r *http.Request) {
	//remove entry from the session manager
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

// create handler for SignIn
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./static/html/home.page.tmpl")

}

// create handler for SignInSubmit
func (app *application) SignInSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

// create handler for ScanQrCode
func (app *application) ScanQrCode(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

func RenderTemplates(w http.ResponseWriter, s string) {
	panic("unimplemented")
}

// create handler for ScanQrCodeSubmit
func (app *application) ScanQrCodeSubmit(w http.ResponseWriter, r *http.Request) {

	RenderTemplates(w, "./ui/static/html/home.page.tmpl")

}

func (app *application) dashboard(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionManager.PopString(r.Context(), "flash")
	//render
	data := &templateData{
		Flash: flash,
	}
	RenderTemplate(w, "index.html", data)
}

// IMPLEMENTING CRUD FOR USERS

// Create User
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	// remove the entry from the session manager
	flash := app.sessionManager.PopString(r.Context(), "flash")
	data := &templateData{
		Flash: flash,
	}
	RenderTemplate(w, "signup.page.tmpl", data)
}

// Read User
