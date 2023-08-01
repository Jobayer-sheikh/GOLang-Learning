package ordermessage

import (
	"database/sql"
	"fiber/learning/database"
	"fiber/learning/handler/smsservice"
	"fiber/learning/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (command *OrderMessageCommand) Handle(ctx *fiber.Ctx) (*response.GenericResponse[int64], error) {
	return command.HandleInternal()
}

func (command *OrderMessageCommand) HandleInternal() (*response.GenericResponse[int64], error) {
	var smsCommand = &smsservice.SmsCommand{
		Contacts: command.Phone,
		Msg:      string(command.Text),
	}

	go func() {
		if success, err := smsCommand.Handle(); err != nil {
			log.Println(fmt.Sprintf("Can't send SMS, response %t", success))
		} else {
			log.Println(fmt.Sprintf("Successfully send SMS, response %t", success))
		}
	}()

	var db = database.GetDbContext()
	var result = command.saveMessage(db)

	return &(*result)[0], nil
}

func (command *OrderMessageCommand) saveMessage(db *database.Database) *[]response.GenericResponse[int64] {
	var sqlStatement = "[pfo].[AddOrderMessage] :OrderId, :CustomerId, :LineOfBusiness, :Phone, :OrderStatus, :Text, :Delivered;"

	var sqlParameter = []any{
		sql.Named("OrderId", command.OrderId),
		sql.Named("CustomerId", command.CustomerId),
		sql.Named("LineOfBusiness", command.LineOfBusiness),
		sql.Named("Phone", command.Phone),
		sql.Named("OrderStatus", command.OrderStatus),
		sql.Named("Text", command.Text),
		sql.Named("Delivered", command.Delivered)}

	result, _ := database.FromStoredProcedure[response.GenericResponse[int64]](
		db,
		sqlStatement,
		sqlParameter,
		func(result *response.GenericResponse[int64]) []any {
			return []any{
				&result.Success,
				&result.Status,
				&result.Response,
				&result.Message,
			}
		})

	return result
}
