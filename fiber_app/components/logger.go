package components

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Logger *log.Logger
	ticker *time.Ticker
)

func makeLogger() *log.Logger {
	// Set the alarm to next day
	y, m, d := time.Now().Date()
	nextDay := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Add(time.Hour * 24)
	ticker.Reset(time.Until(nextDay))

	file, err := os.OpenFile(
		fmt.Sprintf("./log/%d-%d-%d", y, m, d),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0744,
	)
	if err != nil {
		log.Println(err)
	}
	return log.New(file, "", log.LstdFlags)
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
		for t := range ticker.C {
			// Different day allow to create new logger
			if t.Day() != yesterday {
				Logger = makeLogger()
				yesterday = time.Now().Day()
			} else {
				ticker.Reset(time.Second)
			}
		}
	}()
}
