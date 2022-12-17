package components

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var LoggerConfig = logger.Config{
	Format: "${ip} ${ips} ${method} ${status} ${path} ${error}\n",
	Done: func(c *fiber.Ctx, logString []byte) {
		// If you want to use UTF-8, change string to []rune.
		Logger.Print(string(logString))
	},
}
var (
	Logger *log.Logger
	ticker *time.Ticker
	file   *os.File
)

func makeLogger() *log.Logger {
	var err error
	// Set the alarm to next day
	y, m, d := time.Now().Date()
	nextDay := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Add(time.Hour * 24)
	ticker.Reset(time.Until(nextDay))
	// Make log file
	if file != nil {
		file.Close()
	}
	file, err = os.OpenFile(
		fmt.Sprintf("./log/%d-%d-%d.log", y, m, d),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0744,
	)
	if err != nil {
		log.Println(err)
	}
	return log.New(file, "", log.Ltime)
}

func LoggerInit() {
	// Start logger
	ticker = time.NewTicker(time.Second)
	Logger = makeLogger()

	// Logger loop
	// Every day make new logger
	go func() {
		defer ticker.Stop()
		yesterday := time.Now().Day()
		for today := range ticker.C {
			// Different day allow to create new logger
			if today.Day() != yesterday {
				Logger = makeLogger()
				yesterday = time.Now().Day()
			} else {
				ticker.Reset(time.Second)
			}
		}
	}()
}
