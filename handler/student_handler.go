package handler

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/ishtiaqhimel/go-api-server/db"
	"github.com/ishtiaqhimel/go-api-server/model"
	"net/http"
)

func StudentGet(students db.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := students.GetAll()
		json.NewEncoder(w).Encode(data)
	}
}

func StudentPost(students db.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.Student
		json.NewDecoder(r.Body).Decode(&req)
		if err := students.Add(req); err != nil {
			render.Status(r, 409)
			render.Render(w, r, ErrRequest(409, err.Error()))
		} else {
			w.Write([]byte("Student Successfully Added"))
		}
	}
}

func StudentUpdate(students db.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		var req model.Student
		json.NewDecoder(r.Body).Decode(&req)
		if err := students.UpdateById(id, req); err != nil {
			render.Status(r, 404)
			render.Render(w, r, ErrRequest(404, err.Error()))
		} else {
			w.Write([]byte("Student Successfully Updated"))
		}
	}
}

func StudentDelete(students db.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		if err := students.DeleteById(id); err != nil {
			render.Status(r, 404)
			render.Render(w, r, ErrRequest(404, err.Error()))
		} else {
			w.Write([]byte("Student Successfully Deleted"))
		}
	}
}
