package main

import (
	"database/sql"
	"log"
	"net/http"
	application "test_app2/application"
	"test_app2/infra/inmemory"
	"test_app2/service"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
	dsn    = "postgres://root:passwd@localhost:5432/test_app2?sslmode=disable"
)

func main() {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	mux := chi.NewRouter()
	handler := application.NewTodoHandler(service.NewTodoService(inmemory.NewInMemory()))
	// handler = application.NewTodoHandler(service.NewTodoService(store.NewTaskRepository(db)))
	handler.Routes(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server started!")
}
