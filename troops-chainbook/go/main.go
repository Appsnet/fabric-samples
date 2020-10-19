package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"chainbook/conf"
	"chainbook/fabric"
)

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	// Connect to database
	if err := fabric.GetContract(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/policies/", GetAllPolicies)
	app.Get("/policy/:policyno", GetPolicy)
	app.Get("/create/", CreatePolicy)

	log.Fatal(app.Listen(conf.Config("HTTP_PORT")))

}
