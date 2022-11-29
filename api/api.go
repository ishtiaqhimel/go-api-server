package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ishtiaqhimel/go-api-server/auth"
	"github.com/ishtiaqhimel/go-api-server/db"
	"github.com/ishtiaqhimel/go-api-server/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func CallRoutes(username, password, port string) {
	students := db.NewStudent()
	subjects := db.NewSubject()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/api/login", auth.BasicAuth(username, password, handler.LogIn))

	// Student APIs
	r.Get("/api/student", auth.JWTAuth(handler.StudentGet(students)))
	r.Post("/api/student", auth.JWTAuth(handler.StudentPost(students)))
	r.Put("/api/student/{id}", auth.JWTAuth(handler.StudentUpdate(students)))
	r.Delete("/api/student/{id}", auth.JWTAuth(handler.StudentDelete(students)))

	// Subject APIs
	r.Get("/api/subject", auth.JWTAuth(handler.SubjectGet(subjects)))
	r.Post("/api/subject", auth.JWTAuth(handler.SubjectPost(subjects)))
	r.Put("/api/subject/{id}", auth.JWTAuth(handler.SubjectUpdate(subjects)))
	r.Delete("/api/subject/{id}", auth.JWTAuth(handler.SubjectDelete(subjects)))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()
	log.Println("Server Started")
	<-stop
	log.Println("Server Stopped")
}
