package todoController

import (
	todo_service "todo_api/service/todo"
	"todo_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoController struct {
	todoService todo_service.ITodoService
}
type ITodoController interface {
	CreateTodo(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	DeleteTodo(c *fiber.Ctx) error
	GetTodoById(c *fiber.Ctx) error
	GetAllTodos(c *fiber.Ctx) error
}

func NewTodoController(todoService todo_service.ITodoService) ITodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (todoController *TodoController) CreateTodo(c *fiber.Ctx) error {
	
	user := c.Locals("user").(*models.Claim)
	
	data := new(models.CreateTodo)
	
    if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := todoController.todoService.CreateTodo(user.User.Id, data.Content)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
		"result" : result,
	})
}

func (todoController *TodoController) UpdateTodo(c *fiber.Ctx) error {
	
	user := c.Locals("user").(*models.Claim)
	
	data := new(models.CreateTodo)
	
	todo_id, err := uuid.Parse(c.Params("id"))
	
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err = c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ok, err := todoController.todoService.UpdateTodo(user.User.Id, todo_id, data.Content)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var message string

	if ok {
		message = "successful"
	}else {
		message = "unsuccessful"	
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": message,
	})
}
func (todoController *TodoController) DeleteTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.Claim)

	todo_id, err := uuid.Parse(c.Params("id"))
	
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	ok, err := todoController.todoService.DeleteTodo(user.User.Id, todo_id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	var message string

	if ok {
		message = "successful"
	}else {
		message = "unsuccessful"	
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": message,
	})
}

func (todoController *TodoController) GetTodoById(c *fiber.Ctx) error {
	
	user := c.Locals("user").(*models.Claim)
	
	todo_id, err := uuid.Parse(c.Params("id"))
	
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	rows, err := todoController.todoService.GetTodoById(user.User.Id, todo_id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"result": rows,
	})

}

func (todoController *TodoController) GetAllTodos(c *fiber.Ctx) error {
	
	user := c.Locals("user").(*models.Claim)
	
	page := c.Query("page")

	if page == "" {
		page = "0"
	}

	_, err := strconv.Atoi(page)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	intPage, err := strconv.Atoi(page)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	rowsResult, err := todoController.todoService.GetAllTodos(user.User.Id, intPage)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": rowsResult,
	})
}
