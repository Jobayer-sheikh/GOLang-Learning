package ordermessage

import (
	"fiber/learning/entity"
	"github.com/gofiber/fiber/v2"
)

func (command *OrderMessageCommand) Handle(ctx *fiber.Ctx) (*entity.OrderMessage, error) {
	orderMessage := new(entity.OrderMessage)
	return orderMessage, nil
}

//func getMessages(db *database.Database) *[]entity.OrderMessage {
//	var sqlStatement = "GetOrderMessages;"
//
//	result, _ := database.FromSqlRaw[entity.OrderMessage](
//		db,
//		sqlStatement,
//		func(result *entity.OrderMessage) []any {
//			return []any{
//				&result.Id,
//				&result.OrderId,
//				&result.CustomerId,
//				&result.LineOfBusiness,
//				&result.Phone,
//				&result.OrderStatus,
//				&result.Text,
//				&result.Delivered,
//				&result.CreatedAt,
//				&result.UpdatedAt,
//			}
//		})
//
//	return result
//}
