package components

import (
	"github.com/gofiber/fiber/v2"
)

type RestaurantTypes struct {
	RestaurantType string `json:"restaurant_type"`
}

func PostRestaurantTypes(c *fiber.Ctx) error {
	data := RestaurantTypes{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	switch result, err := Postgres.
		Exec("INSERT INTO restaurant_types VALUES ($1) ON CONFLICT (restaurant_type) DO NOTHING", data.RestaurantType); err {
	case nil:
		if _, err := result.RowsAffected(); err != nil {
			return err
		}
	default:
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func GetRestaurantTypes(c *fiber.Ctx) error {
	ret := map[string][]string{}
	switch rows, err := Postgres.
		Query("SELECT * FROM restaurant_types order by restaurant_type"); err {
	case nil:
		tmps, tmp := []string{}, ""
		for rows.Next() {
			if err := rows.Scan(&tmp); err != nil {
				return err
			}
			tmps = append(tmps, tmp)
		}
		ret["restaurant_types"] = tmps
	default:
		return err
	}
	return c.JSON(ret)
}
