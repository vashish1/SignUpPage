package main

import (
	"SignUpPage/database"
	// "SignUpPage/models"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	// "github.com/micro/go-micro/debug/log"
)

var cl *mongo.Collection

func init() {
	cl, _ = database.Createdb()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", signuphandler)
	r.HandleFunc("/confirmation",confirm)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":80", r))
}

func signuphandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		{
			t, err := template.ParseFiles("./signup.html")
			if err != nil {
				log.Fatal("Could not parse template files\n",err)
			}
			er := t.Execute(w, "")
			if er != nil {
				log.Fatal("could not execute the files\n",er)
			}
		}
		log.Print("working")
	case "POST":
		{
			fmt.Println(" lets see if it works ")

			a := r.FormValue("First name")
			b := r.FormValue("Last name")
			c := r.FormValue("Email")
			d := r.FormValue("Password")
			e := r.FormValue("Confirm Password")

			if d == e {
				user := database.NewUser(a, b, c, d)
				database.Insertintouserdb(cl, user)
				http.Redirect(w,r,"/confirmation",302)
			} else {
				http.Redirect(w, r, "/signup", 302)
			}
        
		}
	}
}

func confirm(w http.ResponseWriter,r *http.Request){
	t, err := template.ParseFiles("./confirmation.html")
			if err != nil {
				log.Fatal("Could not parse template files\n",err)
			}
			er := t.Execute(w, "")
			if er != nil {
				log.Fatal("could not execute the files\n",er)
			}
}
