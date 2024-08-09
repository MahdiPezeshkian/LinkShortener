package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MahdiPezeshkian/LinkShortener/internal/endpoints"
	"github.com/MahdiPezeshkian/LinkShortener/internal/repositories"
	"github.com/MahdiPezeshkian/LinkShortener/internal/usecases"
	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	linkRepo := repositories.SQLiteLinkRepository(db)
	linkUsecase := usecases.NewLinkUseCase(linkRepo)
	linkEndpoints := endpoints.NewLinkEndpoints(*linkUsecase)

	http.HandleFunc("/link/get", linkEndpoints.GetLink)
	http.HandleFunc("/link/getpaged", linkEndpoints.GetPagedLinks)
	http.HandleFunc("/link", linkEndpoints.CreateLink)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS links (
        id TEXT PRIMARY KEY,
        is_deleted BOOLEAN,
        is_visibled BOOLEAN,
        original_url TEXT,
        short_url TEXT,
        created_at TIMESTAMP,
        modified_at TIMESTAMP,
        expiration TIMESTAMP,
        clicks INTEGER
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
