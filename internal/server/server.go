package server

import (
	"github.com/JDH-LR-994/learn-go-books/internal/handlers"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Run(addr string) error {
	s.mux.HandleFunc("/books", handlers.BooksHandler)
	return http.ListenAndServe(addr, s.mux)
}
