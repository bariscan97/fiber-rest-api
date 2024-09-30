package authController

import (
	"todo_api/service/user"
	"todo_api/models"
	"os"
	"time"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	IAuthController interface {
		Register(c *fiber.Ctx) error
		Login(c *fiber.Ctx) error
	}
	AuthController struct {
		UserServices user_service.IUserService
	}
)

func NewAuthController(userService user_service.IUserService) IAuthController {
	return &AuthController{
		UserServices: userService,
	}
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	validate := validator.New()

	data := new(models.RegisterUserModel)

	var err error

	if err = c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Struct(data); err != nil {
		return  fiber.NewError(404, err.Error())
	}
	
	bytePassword := []byte(data.Password)
	
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	
	if err != nil {
		return fiber.NewError(401, err.Error())
	}
	
	data.Password = string(hash)
	
	user ,err := controller.UserServices.CreateUser(data)
	if err != nil {
		return fiber.NewError(404, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successful",
		"user" :user,
	})
}

func (controller *AuthController) Login(c *fiber.Ctx) error {

	data := new(models.FetchUserModel)

	var err error

	if err = c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := controller.UserServices.GetUserByEmail(data.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	
	expirationTime := time.Now().Add(72 * time.Hour)
	
	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	
	claims := &models.Claim{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(jwtKey)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    tokenString,
		Expires:  time.Now().Add(72000 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful",
		"access_token" :tokenString,
	})
}
