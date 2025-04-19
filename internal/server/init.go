package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joaocansi/simple-api/internal/config"
	"github.com/joaocansi/simple-api/internal/users"
	"github.com/joaocansi/simple-api/storage"
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

	r := gin.Default()
	s := r.Group("/api/v1")

	users.Setup(s, db)
	r.Run()
}