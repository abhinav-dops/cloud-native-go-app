package main

import (
	"cloud-app-api/internal/handler"
	"cloud-app-api/internal/middleware"
	"cloud-app-api/internal/repository"
	"cloud-app-api/internal/service"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	conn := os.Getenv("DB_CONN")

	var repo *repository.PostgresRepo
	var err error

	for i := 0; i < 10; i++ {
		repo, err = repository.NewPostgresRepo(conn)
		if err == nil {
			break
		}
		log.Println("waiting for database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Databse not ready", err)
	}

	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	svc := service.NewItemService(repo)
	h := handler.NewItemHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/items", h.Items)
	mux.HandleFunc("/health", handler.Health)

	log.Println("Server running on :8081")
	loggedMux := middleware.Logger(mux)
	log.Fatal(http.ListenAndServe(":8081", loggedMux))

}
