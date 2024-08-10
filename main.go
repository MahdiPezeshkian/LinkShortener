package main

import (
	"database/sql"
	"log"

	"github.com/MahdiPezeshkian/LinkShortener/cmd"
	"github.com/MahdiPezeshkian/LinkShortener/internal/endpoints"
	"github.com/MahdiPezeshkian/LinkShortener/internal/repositories"
	"github.com/MahdiPezeshkian/LinkShortener/internal/usecases"
	"github.com/gin-gonic/gin"
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
	
	
	
	r := gin.Default()
	
	r.LoadHTMLGlob("templates/redirect.html*")
	r.GET("/l/:shortURL", linkEndpoints.RedirectToOriginalURL)
	r.GET("/api/li/get/:id", linkEndpoints.GetLink)
	r.GET("/api/li/getpaged", linkEndpoints.GetPagedLinks)
	r.POST("/api/li", linkEndpoints.CreateLink)
	r.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
