package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/VictorOliveiraPy/configs"
	"github.com/VictorOliveiraPy/internal/infra/web/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func createMigrationDatabase() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost,
		configs.DBPort,
		configs.DBUser,
		configs.DBPassword,
		configs.DBName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

}

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost,
		configs.DBPort,
		configs.DBUser,
		configs.DBPassword,
		configs.DBName)

	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	userDb := handlers.NewUserHandler(dbConn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configs.JwtExperesIn))

	r.Post("/users", userDb.CreateUser)
	r.Post("/users/generate_token", userDb.GetJWT)

	http.ListenAndServe(":8000", r)

}
