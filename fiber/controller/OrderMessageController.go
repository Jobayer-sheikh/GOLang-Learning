package controller

import (
	"fiber/learning/entity"
	"fiber/learning/handler"
	"fiber/learning/handler/ordermessage"
	"fiber/learning/response"
	"github.com/gofiber/fiber/v2"
)

func OrderMessageControllerWeb(router fiber.Router) {
	router.Get("/order-message", func(ctx *fiber.Ctx) error {
		var query = new(ordermessage.OrderMessageQuery)
		if err := ctx.QueryParser(query); err != nil {
			var r = err.Error()
			return ctx.JSON(response.Error[string](&r))
		}

		res, _ := query.Handle(ctx)
		return ctx.JSON(response.OK[*[]entity.OrderMessage](&res))
	})

	router.Post("/order-message", func(ctx *fiber.Ctx) error {
		var command = new(ordermessage.OrderMessageCommand)
		if err := ctx.BodyParser(command); err != nil {
			var r = err.Error()
			return ctx.JSON(response.Error[string](&r))
		}

		res, _ := command.Handle(ctx)
		return ctx.JSON(response.OK[*response.GenericResponse[int64]](&res))
	})
}

func OrderMessageControllerApp(router fiber.Router) {
	router.Get("/order-message", handler.OrderMessageSendWebHandler)
}
