package main

import (
	"github.com/JDH-LR-994/learn-go-books/internal/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
