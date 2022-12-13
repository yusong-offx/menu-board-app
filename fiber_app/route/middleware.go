package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MiddleWare(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New(
		logger.Config{
			Format: "[${time}] ${ip} ${ips} ${method} ${status} ${path} ${error}\n",
		},
	))
}
