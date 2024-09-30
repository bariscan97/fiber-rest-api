package routes

import (
	auth"todo_api/routes/authRoute"
	todo"todo_api/routes/todoRoute"
	user"todo_api/routes/userRoute"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Routers struct {
	app *fiber.App
 	pool *pgxpool.Pool
}

func NewRouters(app *fiber.App, pool *pgxpool.Pool) *Routers{
	return &Routers{
		app : app,
		pool: pool,
	}
}

func (routers *Routers) InitRouter() {
	auth.AuthRouter(routers.pool, routers.app)
	todo.TodoRouter(routers.pool, routers.app)
	user.UserRouter(routers.pool, routers.app)
}

func (routers *Routers) Start(Addr string) error{
	return routers.app.Listen(Addr)
}