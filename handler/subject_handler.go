package handler

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/ishtiaqhimel/go-api-server/db"
	"github.com/ishtiaqhimel/go-api-server/model"
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
		if err := subjects.Add(req); err != nil {
			render.Status(r, 409)
			render.Render(w, r, ErrRequest(409, err.Error()))
		} else {
			w.Write([]byte("Subject Successfully Added"))
		}
	}
}

func SubjectUpdate(subjects db.SubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		var req model.Subject
		json.NewDecoder(r.Body).Decode(&req)
		if err := subjects.UpdateById(id, req); err != nil {
			render.Status(r, 404)
			render.Render(w, r, ErrRequest(404, err.Error()))
		} else {
			w.Write([]byte("Subject Successfully Updated"))
		}
	}
}

func SubjectDelete(subjects db.SubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		if err := subjects.DeleteById(id); err != nil {
			render.Status(r, 404)
			render.Render(w, r, ErrRequest(404, err.Error()))
		} else {
			w.Write([]byte("Subject Successfully Deleted"))
		}
	}
}
