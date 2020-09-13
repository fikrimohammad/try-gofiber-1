package main

import (
	"log"

	"github.com/fikrimohammad/Premier-League-DB/configs"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
)

func main() {
	app := fiber.New()
	// app.Settings.Prefork = true

	app.Use(middleware.Recover())
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New())

	app = configs.AppRouter().InitAppRouter(app)
	log.Fatal(app.Listen(3000))
}
