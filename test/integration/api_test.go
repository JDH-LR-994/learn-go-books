package integration_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/JDH-LR-994/learn-go-books/internal/server"
)

func TestBooksAPI(t *testing.T) {
	go func() {
		srv := server.NewServer()
		if err := srv.Run(":8080"); err != nil {
			t.Logf("Ошибка сервера: %v", err)
		}
	}()

	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get("http://localhost:8080/books")
	if err != nil {
		t.Fatalf("Ошибка при запросе: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			t.Errorf("Ошибка при закрытии Body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, получили %d", resp.StatusCode)
	}
}
