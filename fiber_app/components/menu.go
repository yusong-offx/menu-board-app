package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type Menu struct {
	RestaurantID int      `json:"restaurant_id"`
	Name         string   `json:"name"`
	MenuType     string   `json:"menu_type"`
	Images       []string `json:"images"`
	Contents     string   `json:"contents"`
	Allergies    string   `json:"allergies"`
}

func PostMenu(c *fiber.Ctx) error {
	data := Menu{}
	switch result, err := Postgres.
		Exec("INSERT INTO menus (restaurant_id,name,menu_type,images,contents,alleries) VALUES ($1,$2,$3,$4,$5,$6)",
			data.RestaurantID, data.Name, data.MenuType, data.Images, data.Contents, data.Allergies); err {
	case nil:
		if _, err := result.RowsAffected(); err != nil {
			return err
		}
	default:
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

type MenuRetStruct struct {
	ID           int      `json:"id"`
	RestaurantID int      `json:"restaurant_id"`
	Name         string   `json:"name"`
	MenuType     string   `json:"menu_type"`
	Images       []string `json:"images"`
	Contents     string   `json:"contents"`
	Allergies    string   `json:"allergies"`
}

type MenuQueryStruct struct {
	RestaurantID []string `query:"restaurant_id"`
	MenuID       []string `query:"menu_id"`
}

func GetMenu(c *fiber.Ctx) error {
	query := MenuQueryStruct{}
	if err := c.QueryParser(&query); err != nil {
		return err
	}
	if query.RestaurantID != nil {
		ret, tmp := [][]MenuRetStruct{}, MenuRetStruct{}
		for _, userID := range query.RestaurantID {
			switch rows, err := Postgres.Query("SELECT * FROM restaurants WHERE user_id = $1", userID); err {
			case nil:
				byRestaurant := []MenuRetStruct{}
				for rows.Next() {
					if err := rows.Scan(&tmp.ID, &tmp.RestaurantID, &tmp.Name, &tmp.MenuType, pq.Array(&tmp.Images), &tmp.Allergies); err != nil {
						return err
					}
					byRestaurant = append(byRestaurant, tmp)
				}
				ret = append(ret, byRestaurant)
			default:
				return err
			}
		}
		return c.JSON(ret)
	} else {
		// By restaurant_id
		ret, tmp := []MenuRetStruct{}, MenuRetStruct{}
		for _, menuID := range query.MenuID {
			if err := Postgres.QueryRow("SELECT * FROM restaurants WHERE id = $1", menuID).Scan(&tmp.ID, &tmp.RestaurantID, &tmp.Name, &tmp.MenuType, pq.Array(&tmp.Images), &tmp.Allergies); err != nil {
				return err
			}
			ret = append(ret, tmp)
		}
		return c.JSON(ret)
	}
}
