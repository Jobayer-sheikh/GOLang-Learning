package handler

import (
	"fiber/learning/database"
	"fiber/learning/entity"
	"fiber/learning/response"
	"github.com/gofiber/fiber/v2"
)

func OrderMessageSendWebHandler(ctx *fiber.Ctx) error {

	var db = database.GetDbContext()

	messages := getMessages(db)
	return ctx.JSON(response.OK[*[]entity.OrderMessage](&messages))
}

func getMessages(db *database.Database) *[]entity.OrderMessage {
	var sqlStatement = "GetOrderMessages;"

	result, _ := database.FromSqlRaw[entity.OrderMessage](
		db,
		sqlStatement,
		func(result *entity.OrderMessage) []any {
			return []any{
				&result.Id,
				&result.OrderId,
				&result.CustomerId,
				&result.LineOfBusiness,
				&result.Phone,
				&result.OrderStatus,
				&result.Text,
				&result.Delivered,
				&result.CreatedAt,
				&result.UpdatedAt,
			}
		})

	return result
}
