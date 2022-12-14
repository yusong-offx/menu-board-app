package components

import "github.com/gofiber/fiber/v2"

func FuncJwt(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func generateJwt() {

}
