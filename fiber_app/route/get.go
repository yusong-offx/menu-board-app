package route

import (
	"github.com/gofiber/fiber/v2"
	fn "github.com/yusong-offx/menu-board/components"
)

func AllGet(app *fiber.App) {
	// /user/signup
	signup.Get("/:id", fn.PostSignUpUsersChecker)

	// /restaurant
	restaurant.Get("/types", fn.GetRestaurantTypes)
	restaurant.Get("/info", fn.GetRestaurants) //query
	restaurant.Get("/menu", fn.GetMenu)        //query
}
