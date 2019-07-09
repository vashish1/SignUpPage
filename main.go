package main

import "github.com/gorilla/mux"

func main() {
	db,err:=datastore.Newdatastore()
	if err!=nil{
		log.Print(err)
	}
	defer db.close()
	env:=utility.Env{DB:db}

	r := mux.NewRouter()
	r.Handle("/signup",handler.signuphandler(&env)).Methods("GET","POST")
	http.ListenAndServe(":8080",nil)
   
}
