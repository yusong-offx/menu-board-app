package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yusong-offx/menu-board/components"
)

func MiddleWare(app *fiber.App) {
	app.Use(logger.New(
		logger.Config{
			Format: "${ip} ${ips} ${method} ${status} ${path} ${error}\n",
			Done: func(c *fiber.Ctx, logString []byte) {
				// If you want to use UTF-8, change string to []rune.
				components.Logger.Print(string(logString))
			},
		},
	))
	app.Use(recover.New())
}
