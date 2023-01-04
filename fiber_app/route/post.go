package route

import (
	"github.com/gofiber/fiber/v2"
	fn "github.com/yusong-offx/menu-board/components"
)

func AllPost(app *fiber.App) {
	// /user
	user.Post("/login", fn.PostLogin)

	// /user/signup
	signup.Post("/", fn.PostSignUp)

	// /restaurant
	restaurant.Post("/", fn.PostRestaurants)
	restaurant.Post("/types", fn.PostRestaurantTypes)
	restaurant.Post("/menu", fn.PostMenu)
	restaurant.Post("/menu/images", fn.PostMenuImagesUpload)
}
