package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glycek/edu-adapt-api/config"
	// Драйвер для Swagger (пока пустой импорт, чтобы не ругался)
	//_ "github.com/glycek/edu-adapt-api/docs"
)

// @title Edu-Adapt API
// @version 1.0
// @description API для управления учебными материалами
// @host localhost:8080
// @BasePath /
func main() {

	cfg := config.MustLoad()
	// Инициализируем Gin (роутер)
	router := gin.Default()

	// Health-check (чтобы проверить, что всё работает)
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
