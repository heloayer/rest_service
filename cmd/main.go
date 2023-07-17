package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/heloayer/rest/initialize"
	"github.com/heloayer/rest/internal/interfaces"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/my_database")
	if err != nil {
		log.Fatal("error when connecting to database ")
	}
	defer db.Close()

	r := gin.Default()
	interfaces.Routes(r)

	err = initialize.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err := r.Run(); err != nil {
		log.Fatal("error when starting the server:", err)
	}
}
