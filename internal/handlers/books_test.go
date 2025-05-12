package handlers_test

import (
	"encoding/json"
	"github.com/JDH-LR-994/learn-go-books/internal/handlers"
	"github.com/JDH-LR-994/learn-go-books/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPOSTHandler(t *testing.T) {
	jsonStr := `{"title": "Test Book","author": "test"}`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handlers.BooksHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusCreated)
	}

	var book models.Book
	if err := json.NewDecoder(rr.Body).Decode(&book); err != nil {
		t.Fatal("error decoding response body:", err)
	}

	if book.Title != "Test Book" {
		t.Errorf("handler returned wrong title: got %v want %v", book.Title, "Test Book")
	}
}

func TestGETHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rr := httptest.NewRecorder()
	handlers.BooksHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}
}
