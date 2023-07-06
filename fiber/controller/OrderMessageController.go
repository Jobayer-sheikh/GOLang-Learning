package controller

import (
	"fiber/learning/handler"
	"github.com/gofiber/fiber/v2"
)

func OrderMessageControllerWeb(router fiber.Router) {
	router.Get("/order-message", handler.OrderMessageWebHandler)
	router.Post("/order-message", handler.OrderMessageSendWebHandler)
}

func OrderMessageControllerApp(router fiber.Router) {
	router.Get("/order-message", handler.OrderMessageSendWebHandler)
}
