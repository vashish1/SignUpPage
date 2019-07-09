package datastore

import "github/SignUpPage/utility"

//Datastore is an interface for database functionality
type Datastore interface {
	Createuser(u *utility.User) error
	Getuser(username string) (*utility.User, error)
	close()
}

//Newdatastore is a function used to open a mongoDB datastore
func Newdatastore(dbconnectionstring string) (Datastore, error) {
	return NewmongoDB(dbconnectionstring)

}
