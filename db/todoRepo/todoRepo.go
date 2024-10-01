package todoRepo

import (
	"context"
	"fmt"
	"time"
	"todo_api/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TodoRepository struct {
	pool *pgxpool.Pool
}
type ITodoRepository interface {
	CreateTodo(user_id uuid.UUID, content string) (models.FetchTodoModel,error)
	UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) (bool, error)
	DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) (bool, error)
	GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error)
	GetAllTodos(user_id uuid.UUID, page int) ([]models.FetchTodoModel, error)
}

func NewUserRepo(pool *pgxpool.Pool) ITodoRepository {
	return &TodoRepository{
		pool: pool,
	}
}

func (todoRepo *TodoRepository) CreateTodo(user_id uuid.UUID, content string) (models.FetchTodoModel, error) {
	ctx := context.Background()

	sql := `INSERT INTO todos(user_id ,content) VALUES($1 ,$2) RETURNING id, content, user_id, created_at`

	todo := models.FetchTodoModel{}

	err := todoRepo.pool.QueryRow(ctx, sql, user_id, content).Scan(&todo.Id, &todo.Content, &todo.CreateAt)
	
	if err != nil {
		return models.FetchTodoModel{}, fmt.Errorf(err.Error())
	}
    
	return todo, nil
}

func (todoRepo *TodoRepository) UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) (bool, error) {
	ctx := context.Background()

	sql := `
		UPDATE todos
		SET 
			content = $1
		WHERE 
			user_id = $2 AND id = $3

		`
	result, err := todoRepo.pool.Exec(ctx, sql, content, user_id, todo_id)
	
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	
	return result.RowsAffected() == 1, nil 
}

func (todoRepo *TodoRepository) DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) (bool, error) {
	ctx := context.Background()

	sql := `
		DELETE FROM todos
		WHERE user_id = $1 AND id = $2
		`
	result, err := todoRepo.pool.Exec(ctx, sql, user_id, todo_id)
	
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
		
	return result.RowsAffected() == 1, nil 
}

func (todoRepo *TodoRepository) GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error) {
	ctx := context.Background()

	sql := `
		SELECT * FROM todos 
		WHERE user_id = $1 AND id = $2
	`
	rows := todoRepo.pool.QueryRow(ctx, sql, user_id, todo_id)

	var (
		id       uuid.UUID
		userID   uuid.UUID
		content  string
		createAt time.Time
	)

	err := rows.Scan(&id, &content, &userID, &createAt)

	if err != nil {
		return models.FetchTodoModel{}, err
	}

	return models.FetchTodoModel{
		Id:       id,
		Content:  content,
		CreateAt: createAt,
	}, nil
}

func (todoRepo *TodoRepository) GetAllTodos(user_id uuid.UUID, page int) ([]models.FetchTodoModel, error) {
	ctx := context.Background()

	sql := `
		SELECT id, content ,user_id ,created_at FROM todos 
		WHERE user_id = $1
		ORDER BY created_at
		LIMIT 15 
        OFFSET $2 * 15
	`
	rows, err := todoRepo.pool.Query(ctx, sql, user_id, page)
	
	if err != nil {
		return []models.FetchTodoModel{}, err
	}

	var Todos []models.FetchTodoModel

	for rows.Next() {
		var (
			id       uuid.UUID
			content  string
			userID   uuid.UUID
			createAt time.Time
			DbError  error
		)
		DbError = rows.Scan(&id, &content, &userID, &createAt)
		if DbError != nil {
			return []models.FetchTodoModel{}, DbError
		}
		Todos = append(Todos, models.FetchTodoModel{
			Id:       id,
			Content:  content,
			CreateAt: createAt,
		})
	}
	return Todos, nil
}
