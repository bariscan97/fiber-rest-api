package usersController

import (
	user_service "todo_api/service/user"
	"todo_api/models"
	"github.com/gofiber/fiber/v2"
	
)

type UserController struct {
	UserServices user_service.IUserService
}
type IUserController interface {
	GetMe(c *fiber.Ctx) error
	DeleteMe(c *fiber.Ctx) error
	UpdateUsername(c *fiber.Ctx) error
}

func NewUserController(userService user_service.IUserService) IUserController {
	return &UserController{
		UserServices: userService,
	}
}

func (controller *UserController) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.Claim)
	return c.Status(200).JSON(user)

}

func (controller *UserController) DeleteMe(c *fiber.Ctx) error {

	user := c.Locals("user").(*models.Claim)

	err := controller.UserServices.DeleteMe(user.User.Id)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(200).JSON(fiber.Map{
		"delete": "successful",
	})
}

func (controller *UserController) UpdateUsername(c *fiber.Ctx) error {

	user := c.Locals("user").(*models.Claim)
	
	NewUsername := c.Params("username")

	if NewUsername == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username not valid"})
	}

	err := controller.UserServices.UpdateUsername(user.User.Id, NewUsername)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Status(200).JSON(fiber.Map{
		"delete": "successful",
	})
}

