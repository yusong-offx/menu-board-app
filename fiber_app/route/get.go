package route

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/menu-board/components"
)

func AllGet(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		data := struct {
			Id   int               `json:"id"`
			File map[string]string `json:"file"`
		}{}
		var tmp []byte
		if err := components.Postgres.QueryRow("select * from test").Scan(&data.Id, &tmp); err != nil {
			return err
		}
		json.Unmarshal(tmp, &data.File)
		return c.JSON(data)
	})
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello-world!")
	})
	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("error test")
	})
}
