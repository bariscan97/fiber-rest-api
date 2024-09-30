package todoRoute

import (
	controller "todo_api/controller/todoController"
	repository "todo_api/db/todoRepo"
	service "todo_api/service/todo"
	"todo_api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TodoRouter(pool *pgxpool.Pool, router fiber.Router) {
	todoRepository := repository.NewUserRepo(pool)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	todo := router.Group("/todo")
	
	todo.Use(middleware.Auth)

	todo.Get("/", todoController.GetAllTodos)
	todo.Post("/", todoController.CreateTodo)
	todo.Delete("/:id", todoController.DeleteTodo)
	todo.Get("/:id", todoController.GetTodoById)
	todo.Patch("/:id", todoController.UpdateTodo)
}
