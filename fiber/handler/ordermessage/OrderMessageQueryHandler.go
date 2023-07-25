package ordermessage

import (
	"database/sql"
	"fiber/learning/database"
	"fiber/learning/entity"
	"github.com/gofiber/fiber/v2"
)

func (query *OrderMessageQuery) Handle(ctx *fiber.Ctx) (*[]entity.OrderMessage, error) {

	var db = database.GetDbContext()

	messages, err := getMessages(db, query.OrderId)
	return messages, err
}

func getMessages(db *database.Database, orderId int64) (*[]entity.OrderMessage, error) {
	var sqlStatement = "[pfo].GetOrderMessages @OrderId;"
	var sqlParameter = []any{sql.NamedArg{Name: "OrderId", Value: orderId}}

	result, err := database.FromStoredProcedure[entity.OrderMessage](
		db,
		sqlStatement,
		sqlParameter,
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

	return result, err
}
