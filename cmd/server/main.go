package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/VictorOliveiraPy/configs"
	_ "github.com/VictorOliveiraPy/docs"
	"github.com/VictorOliveiraPy/internal/infra/db"
	"github.com/VictorOliveiraPy/internal/infra/logger"
	"github.com/VictorOliveiraPy/internal/infra/web"
	"github.com/VictorOliveiraPy/internal/service"
	"github.com/VictorOliveiraPy/internal/usecase"

	_ "github.com/VictorOliveiraPy/internal/infra/web"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger.Info("About to start user application")

	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost,
		configs.DBPort,
		configs.DBUser,
		configs.DBPassword,
		configs.DBName)

	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}
	defer dbConn.Close()
	createMigrationDatabase(dbConn)

	//entityHandler := handlers.NewEntityHandler(dbConn)
		// Criação do caso de uso CreateUserUseCase

		// Criação do manipulador WebUserHandler

	// Criação da instância do WebUserHandler
	userRepository := db.NewUserRepository(dbConn)
	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)
	userService := service.NewUserService(*createUserUseCase, userRepository)
	webUserHandler := web.NewWebUserHandler(userService)




	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configs.JwtExperesIn))

	

	// r.Route("/gyms", func(r chi.Router) {
	// 	r.Post("/", entityHandler.CreateGym)
	// 	r.Get("/{id}", entityHandler.GetByGym)
	// 	r.Get("/", entityHandler.GetAllGyms)

	// })

	//r.Post("/users", entityHandler.CreateUser)
	//r.Post("/users/generate_token", entityHandler.GetJWT)

	r.Post("/users", webUserHandler.Create)
	//r.Post("/students", entityHandler.Createstudent)

	// Rota de health-check
	r.Get("/health-check", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ping pong"))
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)

}
