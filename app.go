package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	h "github.com/guiceolin/minigo/handlers"
	"github.com/guiceolin/minigo/interactors"
	"github.com/guiceolin/minigo/models"
	"github.com/guiceolin/minigo/repositories/postgres"
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

	urlRepo := postgres.PostgresUrlRepository{DB: db}
	urlInteractor := interactors.UrlInteractor{Repo: urlRepo}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", h.RootHandler())
	r.Post("/urls", h.CreateUrlHandler(urlInteractor))
	r.Get("/urls/new", h.NewURLFormHandler())
	r.Get("/{short}", h.UnshortURLHandler(urlInteractor))
	r.Get("/{short}/info", h.GetShortURLInfo(urlInteractor))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
