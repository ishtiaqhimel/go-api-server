package handler

import (
	"encoding/json"
	"github.com/ishtiaqhimel/slice-tricks/db"
	"net/http"
)

func SubjectGet(students db.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := students.GetAll()
		json.NewEncoder(w).Encode(data)
	}
}
