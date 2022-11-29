package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ishtiaqhimel/go-api-server/auth"
	"github.com/ishtiaqhimel/go-api-server/db"
	"github.com/ishtiaqhimel/go-api-server/handler"
	"github.com/ishtiaqhimel/go-api-server/utils"
	"log"
	"net/http"
)

func main() {
	students := db.NewStudent()
	subjects := db.NewSubject()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", auth.BasicAuth(handler.LogIn))

	// Student APIs
	r.Get("/student", auth.JWTAuth(handler.StudentGet(students)))
	r.Post("/student", handler.StudentPost(students))
	r.Put("/student/{id}", handler.StudentUpdate(students))
	r.Delete("/student/{id}", handler.StudentDelete(students))

	// Subject APIs
	r.Get("/subject", handler.SubjectGet(subjects))
	r.Post("/subject", handler.SubjectPost(subjects))
	r.Put("/subject/{id}", handler.SubjectUpdate(subjects))
	r.Delete("/subject/{id}", handler.SubjectDelete(subjects))

	fmt.Println("Serving on port " + utils.PORT)
	log.Fatal(http.ListenAndServe(utils.PORT, r))
}
