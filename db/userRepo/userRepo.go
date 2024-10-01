package userRepo

import (
	"context"
	"fmt"
	"todo_api/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	UserRepository struct {
		pool *pgxpool.Pool
	}
	IUserRepo interface {
		CreateUser(data *models.RegisterUserModel) (*models.FetchUserModel, error)
		UpdateUsername(id uuid.UUID, username string) (bool, error)
		DeleteMe(id uuid.UUID) (bool, error)
		GetUserByEmail(email string) (*models.FetchUserModel, error)
	}
)

func NewUserRepo(pool *pgxpool.Pool) IUserRepo {
	return &UserRepository{
		pool: pool,
	}
}

func (userRepo *UserRepository) CreateUser(data *models.RegisterUserModel) (*models.FetchUserModel, error) {
	ctx := context.Background()

	sql := `INSERT INTO users(username,email,password) VALUES($1 ,$2, $3) RETURNING username ,email`

	createdUser := &models.FetchUserModel{}

    err := userRepo.pool.QueryRow(ctx, sql, data.Username, data.Email, data.Password).Scan(&createdUser.Username, &createdUser.Email)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return createdUser, nil
}

func (userRepo *UserRepository) GetUserByEmail(email string) (*models.FetchUserModel, error) {

	ctx := context.Background()

	sql := `SELECT id , username, email, password FROM users WHERE email = $1`

	Rows := userRepo.pool.QueryRow(ctx, sql, email)

	user := &models.FetchUserModel{}

	err := Rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return &models.FetchUserModel{}, err
	}

	return user, nil
}

func (userRepo *UserRepository) UpdateUsername(id uuid.UUID, username string) (bool, error) {

	ctx := context.Background()
	
	sql := `
		UPDATE users
		SET
			username = $2
		WHERE id = $1
	`
	result, err := userRepo.pool.Exec(ctx, sql, id, username)

	if err != nil {
		return false ,fmt.Errorf(err.Error())
	}

	return result.RowsAffected() == 1, nil 
}

func (userRepo *UserRepository) DeleteMe(id uuid.UUID) (bool, error) {
	ctx := context.Background()

	sql := `
		DELETE FROM users
		WHERE id = $1
	`
	result, err := userRepo.pool.Exec(ctx, sql, id)

	if err != nil {
		return false ,fmt.Errorf(err.Error())
	}

	return result.RowsAffected() == 1, nil 
}
