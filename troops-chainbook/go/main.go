package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"chainbook/conf"
)

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	// Connect to database
	if err := GetContract(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	app.Post("/api/login", func(c *fiber.Ctx) error {
		return c.SendString("Logined")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/policies/", GetAllPolicies)
	app.Get("/api/policy/:policyno", GetPolicy)

	log.Fatal(app.Listen(conf.Config("HTTP_PORT")))

}
