package handler

import (
	"encoding/json"
	"github.com/ishtiaqhimel/slice-tricks/db"
	"github.com/ishtiaqhimel/slice-tricks/model"
	"net/http"
)

func SubjectGet(subjects db.SubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := subjects.GetAll()
		json.NewEncoder(w).Encode(data)
	}
}

func SubjectPost(subjects db.SubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.Subject
		json.NewDecoder(r.Body).Decode(&req)
		subjects.Add(req)
		w.Write([]byte("Subject Successfully Added"))
	}
}
