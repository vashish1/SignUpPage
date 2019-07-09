package utility

import "time"
//User ......
type User struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Firstname         string `json:"firstname" bson:"firstname"`
	Lastname          string `json:"lastname" bson:"lastname"`
	Email             string `json:"email" bson:"email"`
	PasswordHash      string `json:"password_hash" bson:"password_hash"`
	Timestampcreated  int64  `json:"timestampcreated" bson:"timestampcreated"`
	Timestampmodified int64  `json:"timestampmodified" bson:"timestampmodified"`
}
//Newuser .....
func Newuser(firstname string, lastname string, email string, password string) *User {

	Password := utility.SHA256ofstring(password)
	now := time.Now()
	Unixtimestamp := now.Unix()
	U := User{UUID: utility.GenerateUUID(), Firstname: firstname, Lastname: lastname, Email: email, PasswordHash: Password, Timestampcreated: Unixtimestamp, Timestampmodified: Unixtimestamp}
	return &U
}
