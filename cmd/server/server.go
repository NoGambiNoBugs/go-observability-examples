package main

import "github.com/NoGambiNoBugs/go-observability-examples/internal/bootstrap"

func main() {
	app, err := bootstrap.Setup()
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
