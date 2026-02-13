package main

import (
	"cloud-app-api/internal/handler"
	"cloud-app-api/internal/repository"
	"cloud-app-api/internal/service"
	"log"
	"net/http"
	"os"
)

func main() {

	conn := os.Getenv("DB_CONN")

	repo, err := repository.NewPostgresRepo(conn)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.NewItemService(repo)
	h := handler.NewItemHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/items", h.Items)

	log.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))

}
