package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/menu-board/components"
)

func AllPost(app *fiber.App) {
	signup := app.Group("/signup")
	signup.Post("/", components.FuncSignUp)
	signup.Post("/check/id", components.FuncSignUpUsersChecker)

	app.Post("/login", components.FuncLogin)
}
