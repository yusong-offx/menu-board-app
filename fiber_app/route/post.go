package route

import (
	"github.com/gofiber/fiber/v2"
	fn "github.com/yusong-offx/menu-board/components"
)

func AllPost(app *fiber.App) {
	////////
	// User
	user.Post("/login", fn.PostLogin)

	// /user/signup/*
	signup.Post("/", fn.PostSignUp)
	signup.Post("/check/id", fn.PostSignUpUsersChecker)

	//////////////
	// Restaurant
	restaurant.Post("/type", fn.PostRestaurantTypes)
}
