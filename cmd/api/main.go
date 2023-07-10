package main

import (
	"flag"
	"fmt"
	"go-movies/internal/repository"
	"go-movies/internal/repository/dbrepo"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {

	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")

	flag.Parse()

	conn, err := app.connectToDB()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	if err != nil {
		log.Fatal(err)
	}

	app.Domain = "localhost"

	log.Println("Starting app on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
