package datastore

import (
	"github/SignUpPage/utility"

	"gopkg.in/mgo.v2"
)

//MongoDBdatastore is a session opening function for mongoDB
type MongoDBdatastore struct {
	*mgo.Session
}

//NewmongoDB opens a database
func NewmongoDB(url string) (*MongoDBdatastore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDBdatastore{
		Session: session,
	}, nil

}
//CreateUser creates a user in database
func (m *MongoDBdatastore) CreateUser(user *utility.User) error {
    session:m.Copy()
	defer session.close()
	usercollection := session.DB("SignUpPage").C(user)
	err:=usercollection.Insert(user)
	if err!=nil{
		return err
	}
	return nil
}

//Getuser function gets a user information
func (m *MongoDBdatastore)Getuser(email string) (*utility.User,error){
     session:m.Copy()
	defer session.close()
	usercollection := session.DB("SignUpPage").C(user)
	u:=utility.User{}
	err:=usercollection.find(bson.M{"Email":email}).One(&u)
	if err!=nil{
		return nil,err
	}
	return &u,nil
}
//Close closes the datastore
func  (m *MongoDBdatastore)Close(){
	m.close()
}