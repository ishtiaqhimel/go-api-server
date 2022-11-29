package handler

import (
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type ErrResponse struct {
	HTTPStatusCode int    `json:"code"`
	StatusText     string `json:"message"`
	TimeStamp      string `json:"timeStamp"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRequest(errCode int, msg string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: errCode,
		StatusText:     msg,
		TimeStamp:      time.Now().Format(time.RFC1123),
	}
}
