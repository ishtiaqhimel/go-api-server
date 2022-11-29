package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

var (
	username string
	password string
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username = os.Getenv("username")
		password = os.Getenv("password")

		var data map[string]string
		json.NewDecoder(r.Body).Decode(&data)
		key := data["username"] + ":" + data["password"]
		enKey := "Basic " + base64.StdEncoding.EncodeToString([]byte(key))
		r.Header.Set("Authorization", enKey)
		user, pass, authOk := r.BasicAuth()
		if authOk == false {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		if username != user || password != pass {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func GetToken() (string, error) {
	key := []byte("thisisasecretkeyihavegenerated")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"exp":  time.Now().Add(600 * time.Second).Unix(),
	})
	tokenString, err := token.SignedString(key)
	return tokenString, err
}
