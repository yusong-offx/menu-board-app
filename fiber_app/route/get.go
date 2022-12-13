package route

import "github.com/gofiber/fiber/v2"

func AllGet(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello-world!")
	})
}
