package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/me", me)

	RegisterCustomerRouter(v1)
}

func me(c *fiber.Ctx) error {
	message := "Golang, Kafka, CDC, Worker, Go Fiber, Sql Server e Slack"
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}
