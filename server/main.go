package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

//go:generate weaver generate ./...

func main() {
	if err := weaver.Run(context.Background(), serve); err != nil {
		panic(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	lis weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	app.Logger().Info("Listening on...", "address", app.lis)
	return http.Serve(app.lis, nil)
}
