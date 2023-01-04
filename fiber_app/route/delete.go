package route

import (
	"github.com/gofiber/fiber/v2"
	fn "github.com/yusong-offx/menu-board/components"
)

func AllDelete(app *fiber.App) {
	app.Delete("/static/images/:filename", fn.DeleteMenuImage)
}
