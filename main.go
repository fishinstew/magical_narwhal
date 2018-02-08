package main

import (
	"log"

	"github.com/fishinstew/magical_narwhal/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
