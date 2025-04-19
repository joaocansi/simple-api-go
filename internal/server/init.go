package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaocansi/simple-api/internal/config"
	"github.com/joaocansi/simple-api/internal/storage"
	"github.com/joaocansi/simple-api/internal/users"
)

func Init() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	users.Setup(s, db)
	http.ListenAndServe(":8080", r)
}
