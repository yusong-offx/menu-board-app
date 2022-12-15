package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/menu-board/components"
	"github.com/yusong-offx/menu-board/route"
)

func main() {
	components.PostgresConnect()
	components.RedisConnect()
	components.LoggerInit()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	route.MiddleWare(app)
	route.AllGet(app)
	route.AllPost(app)

	log.Fatal(app.Listen(":3000"))
}
