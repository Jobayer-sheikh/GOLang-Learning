package hook

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool
}

func GitWebHookController(router fiber.Router) {
	router.Post("/notification-service/hook/push", func(ctx *fiber.Ctx) error {
		NotificationServiceHandler(ctx)
		return ctx.JSON(&Response{Success: true})
	})

	router.Post("/premium-fruits/hook/push", func(ctx *fiber.Ctx) error {
		go PremiumFruitHandler(ctx)
		return ctx.JSON(&Response{Success: true})
	})
}
