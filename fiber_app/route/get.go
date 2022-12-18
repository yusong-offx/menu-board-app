package route

import (
	"github.com/gofiber/fiber/v2"
	fn "github.com/yusong-offx/menu-board/components"
)

func AllGet(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello-world!")
	})
	restaurant.Get("/type", fn.GetRestaurantTypes)
}
