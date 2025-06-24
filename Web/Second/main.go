package main

import (
	"log"
	"myapp/database"
	"myapp/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	handlers.SetupRoutes(router, db)

	router.Run(":8080")
}
