package handler

import (
	"encoding/json"
	"github.com/ishtiaqhimel/go-api-server/db"
	"github.com/ishtiaqhimel/go-api-server/model"
	"log"
	"net/http"
	"strings"
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
		students.Add(req)
		w.Write([]byte("Student Successfully Added"))
	}
}

func StudentUpdate(students db.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		var req model.Student
		json.NewDecoder(r.Body).Decode(&req)
		students.UpdateById(id, req)
	}
}

func StudentDelete(students db.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		log.Println(id)
		students.DeleteById(id)
	}
}

func parseURL(url string) string {
	p := strings.Split(url, "/")
	return p[len(p)-1]
}
