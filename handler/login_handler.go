package handler

import (
	"github.com/ishtiaqhimel/go-api-server/auth"
	"net/http"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error generating JWT token: " + err.Error()))
	} else {
		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Token: " + token))
	}
}
