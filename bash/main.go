package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/learning-go-ssh/hook"
	"log"
)

func main() {

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	web := app.Group("/api")

	hook.GitWebHookController(web)

	if err := app.Listen(":6000"); err != nil {
		log.Fatalln(err)
	}

	//app := "echo"
	//
	//arg0 := "-e"
	//arg1 := "Hello world"
	//arg2 := "\n\tfrom"
	//arg3 := "golang"
	//
	//cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	//stdout, err := cmd.Output()
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//// Print the output
	//fmt.Println(string(stdout))
}
