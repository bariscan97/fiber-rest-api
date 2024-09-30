package userroute

import (
	controller "todo_api/controller/usersController"
	repository "todo_api/db/userRepo"
	service "todo_api/service/user"
	"todo_api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func UserRouter(pool *pgxpool.Pool, router fiber.Router) {

	userRepository := repository.NewUserRepo(pool)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	user := router.Group("/user")

	user.Use(middleware.Auth)

	user.Delete("/", userController.DeleteMe)
	user.Get("/me", userController.GetMe)
	user.Patch("/", userController.UpdateUsername)
}