package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/guiceolin/minigo/handlers"
	"github.com/guiceolin/minigo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Url{})

	env := handlers.Env{
		DB: db,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", env.RootHandler)
	r.Post("/urls", env.CreateUrlHandler)
	r.Get("/{short}", env.UnshortURLHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
