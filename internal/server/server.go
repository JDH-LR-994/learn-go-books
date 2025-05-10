package server

import (
	"github.com/JDH-LR-994/learn-go-books/internal/handlers"
	"net/http"
)

func Run() {
	mux := http.NewServeMux()                                         // Создаём "маршрутизатор"
	mux.HandleFunc("/books", http.HandlerFunc(handlers.BooksHandler)) // Подключаем handler
	err := http.ListenAndServe(":8080", mux)                          // Запускаем сервер
	if err != nil {
		return
	}
}
