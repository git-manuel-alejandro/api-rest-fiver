package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Id        string
	FirstName string
	LastName  string
}

func UserHandler(c *fiber.Ctx) error {
	user := User{uuid.NewString(), "Jon", "Doe"}

	return c.Status(fiber.StatusOK).JSON(user)
}

func CreateHandler(c *fiber.Ctx) error {
	user := User{uuid.NewString(), "Jon", "Doe"}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)

}
