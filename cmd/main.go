package main

import (
	"context"
	"log"

	"github.com/mixdjoker/auth/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
