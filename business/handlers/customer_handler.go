package handlers

import (
	"github/jailtonjunior94/go-kafka/business/dtos"
	"github/jailtonjunior94/go-kafka/business/mappings"
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

	response := mappings.ToListResponse(customers)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h CustomerHandler) GetCustomerById(c *fiber.Ctx) error {
	message := "Get Customers By Id"
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func (h CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	request := new(dtos.CustomerRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Unprocessable Entity"})
	}

	if err := request.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	customer := mappings.ToEntity(*request)
	result, err := h.CustomerRepository.Add(&customer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	response := mappings.ToResponse(*result)
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	message := "Update Customer"
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func (h CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNoContent).JSON(nil)
}
