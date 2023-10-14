package v1

import (
	"database/sql"
	"fmt"
	"github.com/dash-sarthak/kanban/internal/database"
	"github.com/dash-sarthak/kanban/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type apiConfig struct {
	DB *database.Queries
}

func connectToDB() (*sql.DB, error) {
	dbURL, envErr := util.LoadFromEnv("DB_URL")
	if envErr != nil {
		log.Fatalf("Error while loading from env: %v", envErr)
	}

	dbConnection, conErr := sql.Open("postgres", dbURL)

	return dbConnection, conErr
}

func createRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	return router
}

func SetupAndRun() {
	port, envErr := util.LoadFromEnv("PORT")
	if envErr != nil {
		log.Fatalf("Error while loading from env: %v", envErr)
	}

	dbConnection, dbErr := connectToDB()
	if dbErr != nil {
		log.Fatalf("Could not connect to DB: %v", dbErr)
	}

	apiCfg := apiConfig{DB: database.New(dbConnection)}

	router := createRouter()
	v1Router := chi.NewRouter()

	v1Router.Post("/author", apiCfg.handleAuthorsCreate)
	v1Router.Get("/authors", apiCfg.handleAuthorsFetch)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%v", port),
	}
	log.Printf("KanBan is now running on port %v\n", port)
	log.Fatal(server.ListenAndServe())
}
