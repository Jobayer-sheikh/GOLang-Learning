package handler

import (
	"fiber/learning/dto"
	"fiber/learning/rabbitmq"
	"fiber/learning/response"
	"github.com/gofiber/fiber/v2"
	"log"
)

func RabbitMQHealthCheckHandler(ctx *fiber.Ctx) error {
	var user = &dto.Admin{
		Id:           1,
		AdminNameEng: "Jobayer",
		AdminNameBng: "Jobayer",
		Phone:        "01752062838",
	}

	var success = false

	if err := rabbitmq.Publish[dto.Admin](user); err == nil {
		success = true
		log.Println("Successfully publish message to RabbitMQ")
	}

	return ctx.JSON(response.OK[bool](&success))
}
