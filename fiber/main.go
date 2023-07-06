package main

import (
	"fiber/learning/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

type User struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	web := app.Group("/api")
	api := app.Group("/api/v1")

	controller.OrderMessageControllerWeb(web)
	controller.OrderMessageControllerApp(api)

	//var res = &User{
	//	Id:    1,
	//	Name:  "Jobayer",
	//	Phone: "01752435262",
	//}
	//
	//var body User
	//var query User
	//
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.JSON(res)
	//})
	//
	//app.Get("/param/:phone", func(c *fiber.Ctx) error {
	//	return c.SendString(c.Params("phone", "none"))
	//})
	//
	//app.Get("/query", func(c *fiber.Ctx) error {
	//	if err := c.QueryParser(&query); err != nil {
	//		return c.JSON(response.ErrorWithMessage[any](nil, err.Error()))
	//	}
	//
	//	return c.JSON(&query)
	//})
	//
	//app.Get("/param-query/:phone", func(c *fiber.Ctx) error {
	//	phone := c.Params("phone", "none")
	//	if err := c.QueryParser(&query); err != nil {
	//		return c.JSON(response.ErrorWithMessage[any](nil, err.Error()))
	//	}
	//
	//	query.Phone = phone
	//	return c.JSON(response.OK[User](&query))
	//})
	//
	//app.Post("/", func(c *fiber.Ctx) error {
	//	if err := c.BodyParser(&body); err != nil {
	//		return err
	//	}
	//
	//	return c.JSON(body)
	//})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
