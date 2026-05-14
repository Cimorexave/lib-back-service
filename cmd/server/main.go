package main

import (
	"log"

	. "github.com/cimorexave/lib-back-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver
)

func main() {
	// Connect to Docker Database
	dsn := "user=user password=password dbname=library_db sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	// Setup Gin Router
	r := gin.Default()

	// Define a Route
	r.GET("/books", func(c *gin.Context) {
		var books []Book
		err := db.Select(&books, "SELECT id, title, author FROM books")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, books)
	})

	// Start Server
	r.Run(":8080")
}
