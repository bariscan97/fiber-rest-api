package models

import (
	"time"

	"github.com/google/uuid"
)

type FetchTodoModel struct {
	Id       uuid.UUID
	Content  string
	CreateAt time.Time
}

type CreateTodo struct {
	Content string
}
