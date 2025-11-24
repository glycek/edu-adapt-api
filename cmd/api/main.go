package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glycek/edu-adapt-api/config"
)

func main() {

	cfg := config.MustLoad()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"db":      cfg.DB.Host, // Просто для теста, что конфиг прочитался
		})
	})

	addr := ":" + cfg.HTTP.Port
	fmt.Printf("App running on %s\n", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
