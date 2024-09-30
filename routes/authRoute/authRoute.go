package authRoute

import (
	controller "todo_api/controller/authController"
	repository "todo_api/db/userRepo"
	service "todo_api/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func AuthRouter(pool *pgxpool.Pool, router fiber.Router) {
	authRepository := repository.NewUserRepo(pool)
	authService := service.NewUserService(authRepository)
	authController := controller.NewAuthController(authService)

	auth := router.Group("/auth")

	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
}




