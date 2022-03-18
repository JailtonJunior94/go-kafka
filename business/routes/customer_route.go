package routes

import (
	"log"

	"github/jailtonjunior94/go-kafka/business/database"
	"github/jailtonjunior94/go-kafka/business/handlers"
	"github/jailtonjunior94/go-kafka/business/repositories"

	"github.com/gofiber/fiber/v2"
)

func RegisterCustomerRouter(router fiber.Router) {
	c, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	r := repositories.NewCustomerRepository(c.Db)
	h := handlers.NewCustomerHandler(r)

	router.Get("/customers", h.GetCustomers)

	router.Get("/customers/:id", h.GetCustomerById)

	router.Post("/customers", h.CreateCustomer)

	router.Put("/customers/:id", h.UpdateCustomer)

	router.Delete("/customers/:id", h.DeleteCustomer)
}
