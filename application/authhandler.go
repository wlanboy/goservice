package application

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

/*authMiddleware*/
func (goservice *GoService) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		realm := goservice.Config.Realm
		userhash := goservice.Config.RealmUser
		secrethash := goservice.Config.RealmSecret
		user, pass, ok := r.BasicAuth()
		if !ok || checkRealmAuthError(user, pass, realm, userhash, secrethash) {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

/*checkRealmAuthError*/
func checkRealmAuthError(user, secret, realm, userhash, secrethash string) bool {
	erru := bcrypt.CompareHashAndPassword([]byte(userhash), []byte(user))
	if erru != nil {
		log.Println(erru)
		return true
	}
	erra := bcrypt.CompareHashAndPassword([]byte(secrethash), []byte(secret))
	if erra != nil {
		log.Println(erra)
		return true
	}

	return false
}
