package main

import (
	"cloud-app-api/internal/handler"
	"cloud-app-api/internal/repository"
	"cloud-app-api/internal/service"
	"log"
	"net/http"
)

func main() {

	repo := repository.NewMemoryRepo()
	svc := service.NewItemService(repo)
	h := handler.NewItemHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/items", h.Items)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
