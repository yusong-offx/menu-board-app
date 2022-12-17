package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yusong-offx/menu-board/components"
)

func MiddleWare(app *fiber.App) {
	app.Use(logger.New(components.LoggerConfig))
	app.Use(recover.New())
}
