package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	dsn := os.Getenv("SPARROW_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")

}
