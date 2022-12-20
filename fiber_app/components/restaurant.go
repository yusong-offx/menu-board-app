package components

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type RestaurantsStruct struct {
	UserID         int               `json:"user_id"`
	Name           string            `json:"name"`
	RestaurantType string            `json:"restaurant_type"`
	Origin         map[string]string `json:"origin"`
	MenuTypes      []string          `json:"menu_types"`
}

func PostRestaurantTypes(c *fiber.Ctx) error {
	var data string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	switch result, err := Postgres.
		Exec("INSERT INTO restaurant_types VALUES ($1) ON CONFLICT (restaurant_type) DO NOTHING", data); err {
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

func PostRestaurants(c *fiber.Ctx) error {
	data := RestaurantsStruct{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	OriginJson, err := json.Marshal(data.Origin)
	if err != nil {
		return err
	}
	switch result, err := Postgres.
		Exec("INSERT INTO restaurants (user_id, name, restaurant_type, origin, menu_types) VALUES ($1,$2,$3,$4,$5)",
			data.UserID, data.Name, data.RestaurantType, OriginJson, pq.Array(data.MenuTypes),
		); err {
	case nil:
		if _, err := result.RowsAffected(); err != nil {
			return err
		}
	default:
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

type GetRestaurantsQueryStruct struct {
	UserID       string   `query:"user_id"`
	RestaurantID []string `query:"restaurant_id"`
}

type RestaurantsRetStruct struct {
	ID             int      `json:"id"`
	UserID         int      `json:"user_id"`
	Name           string   `json:"name"`
	RestaurantType string   `json:"restaurant_type"`
	Origin         string   `json:"origin"`
	MenuTypes      []string `json:"menu_types"`
}

func GetRestaurants(c *fiber.Ctx) error {
	data := GetRestaurantsQueryStruct{}
	if err := c.QueryParser(&data); err != nil {
		return err
	}
	ret, tmp := []RestaurantsRetStruct{}, RestaurantsRetStruct{}
	// By user_id
	if data.UserID != "" {
		switch rows, err := Postgres.Query("SELECT * FROM restaurants WHERE user_id = $1", data.UserID); err {
		case nil:
			for rows.Next() {
				if err := rows.Scan(&tmp.ID, &tmp.UserID, &tmp.Name, &tmp.RestaurantType, &tmp.Origin, &tmp.MenuTypes); err != nil {
					return err
				}
				ret = append(ret, tmp)
			}
		default:
			return err
		}
	} else {
		// By restaurant_id
		for _, restaurantID := range data.RestaurantID {
			if err := Postgres.QueryRow("SELECT * FROM restaurants WHERE id = $1", restaurantID).Scan(&tmp.ID, &tmp.UserID, &tmp.Name, &tmp.RestaurantType, &tmp.Origin, &tmp.MenuTypes); err != nil {
				return err
			}
			ret = append(ret, tmp)
		}
	}
	return c.JSON(ret)
}
