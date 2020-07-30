package main

import (
	"github.com/alexdemen/auto_catalog/router"
	"github.com/alexdemen/auto_catalog/storage"
	"log"
	"net/http"
)

func main() {
	handler := router.NewHandler(storage.NewMemoryStorage())
	log.Fatal(http.ListenAndServe(":9090", handler))
}
