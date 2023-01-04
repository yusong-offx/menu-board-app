package components

import (
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type Menu struct {
	RestaurantID int      `json:"restaurant_id"`
	Name         string   `json:"name"`
	Price        int      `json:"price"`
	MenuType     string   `json:"menu_type"`
	Images       []string `json:"images"`
	Contents     string   `json:"contents"`
	Allergies    string   `json:"allergies"`
}

func PostMenu(c *fiber.Ctx) error {
	data := Menu{}
	switch result, err := Postgres.
		Exec("INSERT INTO menus (restaurant_id,name,price,menu_type,images,contents,alleries) VALUES ($1,$2,$3,$4,$5,$6,$7)",
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
	Price        int      `json:"price"`
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
			if err := Postgres.QueryRow("SELECT * FROM restaurants WHERE id = $1", menuID).Scan(&tmp.ID, &tmp.RestaurantID, &tmp.Name, &tmp.Price, &tmp.MenuType, pq.Array(&tmp.Images), &tmp.Allergies); err != nil {
				return err
			}
			ret = append(ret, tmp)
		}
		return c.JSON(ret)
	}
}

func PostMenuImagesUpload(c *fiber.Ctx) error {
	switch form, err := c.MultipartForm(); err {
	case nil:
		fmt.Println(form.File)
		for _, file := range form.File["image"] {
			if err := c.SaveFile(file, fmt.Sprintf("./assets/images/%s", file.Filename)); err != nil {
				return err
			}
		}
	default:
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func DeleteMenuImage(c *fiber.Ctx) error {
	if filename := c.Params("filename", ""); filename != "" {
		if err := os.Remove("./assets/images/" + filename + ".png"); err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}
	return errors.New("file does not exist")
}
