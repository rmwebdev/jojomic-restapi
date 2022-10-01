package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rmwebdev/jojomic-restapi/app/config"
	"github.com/rmwebdev/jojomic-restapi/app/lib"
	"github.com/rmwebdev/jojomic-restapi/app/routes"
	"github.com/spf13/viper"
)

func init() {
	lib.LoadEnvironment(config.Environment)
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: viper.GetString("PREFORK") == "true",
	})

	routes.Handle(app)
	log.Fatal(app.Listen(":" + viper.GetString("PORT")))
}
