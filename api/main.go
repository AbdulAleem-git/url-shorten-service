package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abdulaleem-git/shorten-url/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/:url", routes.ShortenURL)

}
func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
