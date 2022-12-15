package components

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type SignUpStruct struct {
	LoginID   string `json:"id"`
	Password  []byte `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LoginStruct struct {
	LoginID  string `json:"id"`
	Password []byte `json:"password"`
}

func FuncSignUp(c *fiber.Ctx) error {
	data := SignUpStruct{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if err := Postgres.QueryRow("SELECT login_id FROM users WHERE login_id = $1", data.LoginID).Scan(&data.LoginID); err == nil {
		return errors.New("user_id is already exist")
	}
	hashPassword, err := bcrypt.GenerateFromPassword(data.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	switch result, err := Postgres.Exec("INSERT INTO users (login_id, login_password, first_name, last_name, email) VALUES ($1, $2, $3, $4, $5)", data.LoginID, hashPassword, data.FirstName, data.LastName, data.Email); err {
	case nil:
		n, err := result.RowsAffected()
		switch {
		case err != nil:
			return err
		case n != 1:
			return errors.New("sql insert error")
		}
	default:
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func FuncLogin(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
