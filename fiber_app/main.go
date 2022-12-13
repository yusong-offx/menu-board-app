package main

import (
	"time"

	"github.com/yusong-offx/menu-board/components"
)

func main() {
	// app := fiber.New(fiber.Config{
	// 	Prefork: false,
	// })

	// route.MiddleWare(app)
	// route.AllGet(app)

	// log.Fatal(app.Listen(":3000"))

	components.LoggerInit()
	time.Sleep(time.Second * 11)
}
