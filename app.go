package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/guiceolin/minigo/handlers"
	"github.com/guiceolin/minigo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=minigo sslmode=disable ")
	//db, err := gorm.Open("postgres", "postgresql://postgres@localhost:5432/minigo")
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
