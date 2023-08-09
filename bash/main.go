package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/learning-go-ssh/hook"
	"log"
)

func main() {

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	web := app.Group("/api")

	hook.GitWebHookController(web)

	if err := app.Listen(":6000"); err != nil {
		log.Fatalln(err)
	}
}
