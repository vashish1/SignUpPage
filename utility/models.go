package utility

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/google/uuid"
)

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
	return string(sd)

}
//Env ........
type Env struct{
	DB datastore.Datastore
}