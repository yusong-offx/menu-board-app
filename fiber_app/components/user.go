package components

import "github.com/gofiber/fiber/v2"

type SignUpStruct struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginStruct struct {
	Id       string   `json:"id"`
	Password []byte   `json:"password"`
	Stores   []string `json:"stores"`
}

func FuncSignUp(c *fiber.Ctx) error {
	data := SignUpStruct{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func FuncLogin(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
