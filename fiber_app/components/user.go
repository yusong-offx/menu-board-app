package components

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserStruct struct {
	LoginID   string `json:"login_id"`
	Password  string `json:"lrogin_password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LoginStruct struct {
	LoginID  string `json:"login_id"`
	Password string `json:"login_password"`
}

type LoginRetStruct struct {
	ID        int    `json:"id"`
	LoginID   string `json:"login_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func PostSignUpUsersChecker(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return errors.New("no param")
	}
	// Check user already exist
	if err := Postgres.QueryRow(
		"SELECT login_id FROM users WHERE login_id = $1", id).
		Scan(&id); err == nil {
		return c.Status(fiber.StatusConflict).SendString("user_id is already exist")
	}
	return c.SendStatus(fiber.StatusOK)
}

func PostSignUp(c *fiber.Ctx) error {
	data := UserStruct{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// Bcrypt password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Insert into database
	switch result, err := Postgres.Exec(
		"INSERT INTO users (login_id, login_password, first_name, last_name, email) VALUES ($1, $2, $3, $4, $5)",
		data.LoginID, hashPassword, data.FirstName, data.LastName, data.Email); err {
	case nil:
		if _, err := result.RowsAffected(); err != nil {
			return err
		}
	default:
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func PostLogin(c *fiber.Ctx) error {
	var hashPassword []byte
	data, ret := LoginStruct{}, LoginRetStruct{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if err := Postgres.QueryRow("SELECT * FROM users WHERE login_id=$1", data.LoginID).
		Scan(&ret.ID, &ret.LoginID, &hashPassword, &ret.FirstName, &ret.LastName, &ret.Email); err != nil {
		return c.Status(fiber.StatusConflict).SendString("wrong id")
	}
	if err := bcrypt.CompareHashAndPassword(hashPassword, []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusConflict).SendString("wrong password")
	}
	// JWT 생성
	return c.JSON(ret)
}
