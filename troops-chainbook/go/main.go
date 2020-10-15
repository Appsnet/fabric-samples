/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	// Connect to database
	if err := GetContract(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/policies/", GetAllPolicies)
	app.Get("/policy/:policyno", GetPolicy)

	log.Fatal(app.Listen(Config("HTTP_PORT")))

}
