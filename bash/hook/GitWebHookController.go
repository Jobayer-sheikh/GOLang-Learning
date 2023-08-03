package hook

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool
}

func GitWebHookController(router fiber.Router) {
	router.Get("/hook/push", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&Response{Success: true})
	})

	router.Post("/hook/push", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&Response{Success: true})
	})
}
