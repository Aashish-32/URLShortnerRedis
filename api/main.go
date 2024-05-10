package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aashish-32/URL-Shortener/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL) // " : " indicates that url is a parameter in the route pattern.
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("cant load enviroment variables")
	}

	app := fiber.New()
	port := os.Getenv("PORT")
	app.Use(logger.New())
	SetupRoutes(app)
	log.Fatal(app.Listen(":" + port))

}
