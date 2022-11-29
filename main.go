package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ishtiaqhimel/slice-tricks/db"
	"github.com/ishtiaqhimel/slice-tricks/handler"
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	students := db.NewStudent()
	subjects := db.NewSubject()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Student APIs
	r.Get("/student", handler.StudentGet(students))
	r.Post("/student", handler.StudentPost(students))
	r.Put("/student/{id}", handler.StudentUpdate(students))
	r.Delete("/student/{id}", handler.StudentDelete(students))

	// Subject APIs
	r.Get("/subject", handler.SubjectGet(subjects))
	r.Post("/subject", handler.SubjectPost(subjects))

	fmt.Println("Serving on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
