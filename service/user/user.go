package user_service

import (
	"todo_api/db/userRepo"
	"todo_api/models"

	"github.com/google/uuid"
)

type (
	IUserService interface {
		CreateUser(data *models.RegisterUserModel) (*models.FetchUserModel, error)
		GetUserByEmail(email string) (*models.FetchUserModel, error)
		UpdateUsername(id uuid.UUID, username string) error
		DeleteMe(id uuid.UUID) error
	}
	UserService struct {
		userRepo userRepo.IUserRepo
	}
)

func NewUserService(repository userRepo.IUserRepo) IUserService {
	return &UserService{
		userRepo: repository,
	}
}

func (userService *UserService) CreateUser(data *models.RegisterUserModel) (*models.FetchUserModel, error) {
	return  userService.userRepo.CreateUser(data)
}

func (userService *UserService) GetUserByEmail(email string) (*models.FetchUserModel, error) {
	return userService.userRepo.GetUserByEmail(email)
}

func (userService *UserService) UpdateUsername(id uuid.UUID, username string) error {
	return userService.userRepo.UpdateUsername(id, username)
}

func (userService *UserService) DeleteMe(id uuid.UUID) error {
	return userService.userRepo.DeleteMe(id)
}
