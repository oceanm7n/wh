// Hello world server on port 8080

package main

import (
	"log"

	"github.com/oceanm7n/wh/config"
	"github.com/oceanm7n/wh/server"

	"github.com/spf13/viper"
)

// Function starts the server
func main() {
	// load config from init.go
	if err := config.Init(); err != nil {
		log.Fatalf("Error loading config: %s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}

}
