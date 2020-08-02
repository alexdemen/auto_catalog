package main

import (
	"github.com/alexdemen/auto_catalog/router"
	"github.com/alexdemen/auto_catalog/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handler := router.NewHandler(storage.NewMemoryStorage())
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
