package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

//go:generate weaver generate ./...

func main() {
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
}

func run(ctx context.Context, app *app) error {
	app.Logger().Info("Hello, World!")
	return nil
}
