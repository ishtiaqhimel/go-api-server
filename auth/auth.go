package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/ishtiaqhimel/go-api-server/utils"
	"log"
	"net/http"
	"os"
	"strings"
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
			http.Error(w, "Not Authorized: Username or Password Mismatched", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func GetToken() (string, error) {
	key := []byte(utils.SECRET_KEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"exp":  time.Now().Add(600 * time.Second).Unix(),
	})
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(utils.SECRET_KEY)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		log.Println("JWT Token: ", tokenString)
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Token"))
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		name := claims.(jwt.MapClaims)["name"].(string)
		r.Header.Set("name", name)

		next.ServeHTTP(w, r)
	}
}
