package server

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *Server) MapCustomMiddlewares() {
	s.Fiber.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":   "pong",
			"timestamp": time.Now(),
		})
	})

	s.Fiber.Use(
		CustomRecover(),
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Content-Type, Authorization, Origin, X-Custom-Header",
			AllowMethods: "POST,GET,PUT,DELETE,HEAD,OPTIONS,PATCH",
		}),
	)

	s.Fiber.Get("/panic", func(_ *fiber.Ctx) error {
		panic("this is a test panic")
	})

	s.Fiber.Static("/uploads", "./uploads", fiber.Static{
		Browse: true,
	})
}

func CustomRecover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Request failed: %v", err)

				if jsonErr := c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Something went wrong, please try again later",
				}); jsonErr != nil {
					log.Printf("error sending response: %v, URL: %v", err, c.OriginalURL())
				}
			}
		}()

		return c.Next()
	}
}
