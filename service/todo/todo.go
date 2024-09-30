package todo_service

import (
	"todo_api/db/todoRepo"
	"todo_api/models"

	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo todoRepo.ITodoRepository
}
type ITodoService interface {
	CreateTodo(user_id uuid.UUID, content string) (models.FetchTodoModel,error)
	UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) (bool, error)
	DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) (bool, error)
	GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error)
	GetAllTodos(user_id uuid.UUID, page int) ([]models.FetchTodoModel, error)
}

func NewTodoService(todoRepo todoRepo.ITodoRepository) ITodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (todoService *TodoService) CreateTodo(user_id uuid.UUID, content string) (models.FetchTodoModel,error) {
	return todoService.todoRepo.CreateTodo(user_id, content)
}

func (todoService *TodoService) UpdateTodo(user_id uuid.UUID, todo_id uuid.UUID, content string) (bool, error) {
	return todoService.todoRepo.UpdateTodo(user_id, todo_id, content)

	
}
func (todoService *TodoService) DeleteTodo(user_id uuid.UUID, todo_id uuid.UUID) (bool, error) {
	return todoService.todoRepo.DeleteTodo(user_id, todo_id)
}
func (todoService *TodoService) GetTodoById(user_id uuid.UUID, todo_id uuid.UUID) (models.FetchTodoModel, error) {
	return todoService.todoRepo.GetTodoById(user_id, todo_id)

}

func (todoService *TodoService) GetAllTodos(user_id uuid.UUID, page int) ([]models.FetchTodoModel, error) {
	return todoService.todoRepo.GetAllTodos(user_id, page)
	
}
