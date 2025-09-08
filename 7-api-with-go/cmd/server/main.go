package main

import (
	"mobin.dev/internal/app"
	"mobin.dev/pkg/config"
)

func main() {
	conf := config.Load()
	app.RunServer(conf)
}
