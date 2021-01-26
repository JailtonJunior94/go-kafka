package handlers

import (
	"github/jailtonjunior94/go-kafka/business/repositories"

	"github.com/gofiber/fiber/v2"
)

type ICustomerHandler interface {
	GetCustomers(c *fiber.Ctx) error
	GetCustomerById(c *fiber.Ctx) error
	CreateCustomer(c *fiber.Ctx) error
	UpdateCustomer(c *fiber.Ctx) error
	DeleteCustomer(c *fiber.Ctx) error
}

type CustomerHandler struct {
	CustomerRepository repositories.ICustomerReposity
}

func NewCustomerHandler(repository repositories.ICustomerReposity) ICustomerHandler {
	return &CustomerHandler{CustomerRepository: repository}
}

func (h CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.CustomerRepository.Get()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}

func (h CustomerHandler) GetCustomerById(c *fiber.Ctx) error {
	message := "Get Customers By Id"
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func (h CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	message := "Create Customer"
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": message})
}

func (h CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	message := "Update Customer"
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func (h CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNoContent).JSON(nil)
}
