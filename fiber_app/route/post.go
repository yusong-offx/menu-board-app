package route

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/menu-board/components"
)

func AllPost(app *fiber.App) {
	app.Post("/post", func(c *fiber.Ctx) error {
		data := map[string]string{}
		data["hello"] = "world"
		json_data, err := json.Marshal(data)
		if err != nil {
			return err
		}
		switch result, err := components.Postgres.Exec("insert into test (file) values ($1)", json_data); err {
		case nil:
			n, err := result.RowsAffected()
			if n != 1 || err != nil {
				return nil
			}
		default:
			return err
		}
		return c.SendStatus(fiber.StatusCreated)
	})
}
