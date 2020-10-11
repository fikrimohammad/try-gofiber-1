package main

import (
	"log"

	"github.com/fikrimohammad/Premier-League-DB/configs"
	"github.com/fikrimohammad/Premier-League-DB/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	// app.Settings.Prefork = true

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())

	database.Connect()

	router := configs.AppRouter()
	router.RegisterAPI(app)

	log.Fatal(app.Listen(":3000"))
}
