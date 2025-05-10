package handlers

import (
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
}
