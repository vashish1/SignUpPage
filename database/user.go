package database

import (
	"crypto/sha1"
	"encoding/hex"
	// "time"
	// "SignUpPage/models"

	"github.com/google/uuid"
)

type User struct {
	UUID      string
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

//Newuser .....
func NewUser(first, last, email, pass string) User {

	Password := SHA256ofstring(pass)
	U := User{
		UUID:      GenerateUUID(),
		Firstname: first,
		Lastname:  last,
		Email:     email,
		Password:  Password,
	}
	return U
}

//SHA256ofstring is a function which takes a string a reurns its sha256 hashed form
func SHA256ofstring(p string) string {
	h := sha1.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

//GenerateUUID generates a unique id for every user.
func GenerateUUID() string {

	sd := uuid.New()
	return (sd.String())

}
