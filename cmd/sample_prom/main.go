package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New() 

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello World")
    })
    
    app.Use(logger.New())

    app.Listen(":3000")
}
