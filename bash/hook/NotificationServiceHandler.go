package hook

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os/exec"
)

func NotificationServiceHandler(ctx *fiber.Ctx) {

	var path = "cd /home/raw-application/Notification-Service/"
	run(path)

	var git = "git pull https://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@github.com/Jobayer-sheikh/Notification-Service.git"
	run(git)
}

func run(c string) {
	var (
		err    error
		stdout []byte
	)
	var cmd = exec.Command(c)

	if stdout, err = cmd.Output(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(stdout))
	}
}
