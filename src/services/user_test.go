package services_test

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/models"
	"ecommerce-backend/src/repositories"
	"ecommerce-backend/src/services"
	"ecommerce-backend/src/utils"
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	configs.InitConfigMock()
}

var userCreateMock1 = models.User{
	Provider:  models.LocalProvider,
	Email:     "boat@gmail.com",
	Password:  "1234",
	Firstname: "boat",
	Lastname:  "asdas",
}

func getUserCreatedMock1() models.User {
	u := models.NewUser(userCreateMock1)
	u.ID = primitive.NewObjectID()
	return u
}

func TestCreateUser(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		m := getUserCreatedMock1()
		userRepo.On("Create").Return(&m, nil)
		userRepo.On("FindByEmail").Return(nil, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, _ := userService.CreateUser(userCreateMock1)
		// assert
		assert.NotEmpty(t, result)
		assert.Equal(t, m, *result)
	})

	t.Run("email already exist", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		m := getUserCreatedMock1()
		userRepo.On("FindByEmail").Return(&m, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.CreateUser(userCreateMock1)
		// assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, err.(*fiber.Error))
		assert.Equal(t, 422, err.(*fiber.Error).Code)
	})

	t.Run("repo: find by email error", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		userRepo.On("FindByEmail").Return(nil, errors.New("repo error"))
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.CreateUser(userCreateMock1)
		// assert
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("repo: create error", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		userRepo.On("FindByEmail").Return(nil, nil)
		userRepo.On("Create").Return(nil, errors.New("repo error"))
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.CreateUser(userCreateMock1)
		// assert
		assert.Nil(t, result)
		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	loginDtoMock := models.LoginUserDto{
		Email:    "boat@gmail.com",
		Password: "1234",
	}
	t.Run("repo: find by email error", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		userRepo.On("FindByEmail").Return(nil, errors.New("repo error"))
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.LoginUser(loginDtoMock)
		// assert
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("email not found", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		userRepo.On("FindByEmail").Return(nil, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.LoginUser(loginDtoMock)
		// assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, err.(*fiber.Error))
		assert.Equal(t, 401, err.(*fiber.Error).Code)
	})

	t.Run("password incorrect", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		m := getUserCreatedMock1()
		m.Password = "nowayja"
		hash, _ := utils.Hash(m.Password)
		m.Password = hash
		userRepo.On("FindByEmail").Return(&m, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.LoginUser(loginDtoMock)
		// assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, err.(*fiber.Error))
		assert.Equal(t, 401, err.(*fiber.Error).Code)
	})

	t.Run("login success", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		m := getUserCreatedMock1()
		hash, _ := utils.Hash(m.Password)
		m.Password = hash
		userRepo.On("FindByEmail").Return(&m, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, _ := userService.LoginUser(loginDtoMock)
		// assert
		assert.NotEmpty(t, result)
		assert.IsType(t, &models.Token{}, result)
	})
}
