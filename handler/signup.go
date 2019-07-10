package handler

import (
	"github/SignUpPage/utility"
	"log"
	"net/http"
)

type signup struct {
	Fieldnames []string
	fields     map[string]string
	errors     map[string]string
}

func displaySignupForm(w http.ResponseWriter, r *http.Request, s *signup) {
	utility.RenderTemplate(w, "C:/Go/src/github/SignUpPage/signup.html", s)
}

func displayConfirmation(w http.ResponseWriter, r *http.Request, s *signup) {
	utility.RenderUnsafeTemplate(w, "", s)
}

func populateField(r *http.Request, s *signup) {
	for _, fieldname := range s.Fieldnames {
		s.fields[fieldname] = r.FormValue(fieldname)
	}
}
func validatesignupform(w http.ResponseWriter, r *http.Request, s *signup, e *utility.Env) {
	populateField(r, s)
	if r.FormValue("Firstname") == "" {
		s.errors["Firstname"] = "The form field is required"
	}
	if r.FormValue("Lastname") == "" {
		s.errors["Lastname"] = "The form field is required"
	}
	if r.FormValue("Email") == "" {
		s.errors["Email"] = "The form field is required"
	}
	if r.FormValue("Password") != r.FormValue("Confirm Password") {
		s.errors["Confirm Password"] = "The password and confirm password do not match"
	}

	if len(s.errors) > 0 {
		displaySignupForm(w, r, s)
	} else {
		processSignupForm(w, r, s, e)
	}
}
func processSignupForm(w http.ResponseWriter, r *http.Request, s *signup, e *utility.Env) {
	u := utility.Newuser(r.FormValue("Firstname"), r.FormValue("Lastname"), r.FormValue("Email"), r.FormValue("Password"))
	err := e.DB.Createuser(u)
	if err != nil {
		log.Print(err)

	}
	displayConfirmation(w, r, s)

}

func signuphandler(e *utility.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := signup{}
		s.Fieldnames = []string{"Firstname", "Lastname", "Email"}
		s.fields = make(map[string]string)
		s.errors = make(map[string]string)

		switch r.Method {
		case "GET":
			displaySignupForm(w, r, &s)
		case "POST":
			validatesignupform(w, r, &s, e)
		default:
			displaySignupForm(w, r, &s)
		}

	})
}
