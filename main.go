package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MahdiPezeshkian/LinkShortener/cmd"
	"github.com/MahdiPezeshkian/LinkShortener/internal/endpoints"
	"github.com/MahdiPezeshkian/LinkShortener/internal/repositories"
	"github.com/MahdiPezeshkian/LinkShortener/internal/usecases"
	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db, err := sql.Open("sqlite3", "./LinkService.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := cmd.CreateTable(db); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	if err := cmd.SeedData(db); err != nil {
		log.Fatalf("Error seeding data: %v", err)
	}

	linkRepo := repositories.SQLiteLinkRepository(db)
	linkUsecase := usecases.NewLinkUseCase(linkRepo)
	linkEndpoints := endpoints.NewLinkEndpoints(*linkUsecase)

	http.HandleFunc("/link/get", linkEndpoints.GetLink)
	http.HandleFunc("/link/getpaged", linkEndpoints.GetPagedLinks)
	http.HandleFunc("/link", linkEndpoints.CreateLink)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
