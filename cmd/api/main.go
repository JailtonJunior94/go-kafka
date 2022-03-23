package main

import (
	"fmt"
	"log"

	"github/jailtonjunior94/go-kafka/business/environments"
	"github/jailtonjunior94/go-kafka/business/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	environments.NewConfig()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	routes.RegisterRoutes(app)

	fmt.Printf("🚀 API is running on http://localhost:%d", environments.Port)
	log.Fatal(app.Listen(fmt.Sprintf(":%v", environments.Port)))
}
