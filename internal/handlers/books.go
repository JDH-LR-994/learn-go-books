package handlers

import (
	"crypto/rand"
	"encoding/json"
	"github.com/JDH-LR-994/learn-go-books/models"
	"net/http"
)

// Обработчик для запросов /books
func BooksHandler(resp http.ResponseWriter, req *http.Request) {
	//Обработка GET - запроса
	if req.Method == "GET" {
		books := []models.Book{}
		jsonBytes, err := json.Marshal(books) // Преобразуем данные в JSON
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}
		resp.WriteHeader(http.StatusOK)
		resp.Header().Set("Content-Type", "application/json")
		resp.Write(jsonBytes) // Отправляем JSON
	}
	// Обработчик POST - запроса
	if req.Method == "POST" {
		var book models.Book
		if err := json.NewDecoder(req.Body).Decode(&book); err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}
		if book.Title == "" || book.Author == "" {
			http.Error(resp, "Missing title/author", http.StatusBadRequest)
			return
		}
		book.ID = rand.Text()
		jsonBytes, err := json.Marshal(book)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
		}
		resp.WriteHeader(http.StatusCreated)
		resp.Header().Set("Content-Type", "application/json")
		resp.Write(jsonBytes)
	}
}
