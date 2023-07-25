package controller

import (
	"fiber/learning/handler"
	"github.com/gofiber/fiber/v2"
)

func RabbitMQControllerWeb(router fiber.Router) {
	router.Get("/_health", handler.RabbitMQHealthCheckHandler)
}
