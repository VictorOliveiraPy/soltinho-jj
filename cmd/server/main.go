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
	"github.com/go-chi/jwtauth"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	httpSwagger "github.com/swaggo/http-swagger"
)

func createMigrationDatabase(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
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

// @title           Soltinho JJ
// @version         1.0
// @description     Soltinho API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Victor Hugo
// @contact.url    https://www.linkedin.com/in/victor-hugo-3548a915a/
// @contact.email  oliveiravictordev@gmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	gymDb := handlers.NewGymHandler(dbConn)
	studentdb := handlers.NewStudentHandler(dbConn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configs.JwtExperesIn))


	r.Route("/users", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/{id}", userDb.GetUserFullProfile)
	})

	r.Post("/users", userDb.CreateUser)
	r.Post("/users/generate_token", userDb.GetJWT)

	r.Post("/students", studentdb.Createstudent)

	r.Post("/gyms", gymDb.CreateGym)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)

}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
