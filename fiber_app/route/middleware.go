package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yusong-offx/menu-board/components"
)

var (
	user       fiber.Router
	signup     fiber.Router
	restaurant fiber.Router
)

func MiddleWare(app *fiber.App) {
	user = app.Group("/user")
	signup = user.Group("/signup")
	restaurant = app.Group("/restaurant")

	app.Use(logger.New(components.LoggerConfig))
	app.Use(recover.New())
}
