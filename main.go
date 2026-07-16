package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/alfredoprograma/mks/internal/config"
	"github.com/alfredoprograma/mks/internal/database"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = database.Connect(context.Background(), cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	})

	log.Println(fmt.Sprintf("starting server at %d", cfg.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
