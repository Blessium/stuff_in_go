package main

import (
	"github.com/blessium/metricsgo/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func main() {
	app := fiber.New()

	reg := internal.HandlerRandomMetrics()
	prometheus.Register(reg)

	p := fasthttpadaptor.NewFastHTTPHandler(
		promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
				Registry:          reg,
			},
		),
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/metrics", func(c *fiber.Ctx) error {
		p(c.Context())
		return nil
	})

	app.Use(logger.New())

	app.Listen(":3000")
}
