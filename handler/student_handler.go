package handler

import (
	"encoding/json"
	"github.com/ishtiaqhimel/slice-tricks/db"
	"github.com/ishtiaqhimel/slice-tricks/model"
	"log"
	"net/http"
	"strings"
)

func StudentGet(students db.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := students.GetAll()
		json.NewEncoder(w).Encode(data)
	}
}

func StudentPost(students db.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.Student
		json.NewDecoder(r.Body).Decode(&req)
		students.Add(req)
		w.Write([]byte("Student Successfully Added"))
	}
}

func StudentUpdate(students db.Update) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := parseURL(r.URL.Path)
		var req model.Student
		json.NewDecoder(r.Body).Decode(&req)
		students.UpdateById(id, req)
	}
}

func StudentDelete(students db.Delete) http.HandlerFunc {
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
