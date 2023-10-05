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

var userCreatedMock1 = models.NewUser(userCreateMock1)

func TestCreateUser(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "dsfwdhuasjkdasdsakjdhak"
		userRepo.On("Create").Return(&userCreatedMock1, nil)
		userRepo.On("FindByEmail").Return(nil, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, _ := userService.CreateUser(userCreateMock1)
		// assert
		assert.NotEmpty(t, result)
		assert.Equal(t, userCreatedMock1, *result)
	})

	t.Run("email already exist", func(t *testing.T) {
		// mock
		userRepo := repositories.NewUserRepositoryMock()
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "dsfwdhuasjkdasdsakjdhak"
		userRepo.On("FindByEmail").Return(&userCreatedMock1, nil)
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
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "dsfwdhuasjkdasdsakjdhak"
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
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "dsfwdhuasjkdasdsakjdhak"
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
	t.Run("repo: find by email error", func(t *testing.T) {
		// mock
		loginDto := models.LoginUserDto{
			Email:    "boat@gmail.com",
			Password: "1234",
		}
		userRepo := repositories.NewUserRepositoryMock()
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "dsfwdhuasjkdasdsakjdhak"
		userRepo.On("FindByEmail").Return(nil, errors.New("repo error"))
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.LoginUser(loginDto)
		// assert
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("email not found", func(t *testing.T) {
		// mock
		loginDto := models.LoginUserDto{
			Email:    "boat@gmail.com",
			Password: "1234",
		}
		userRepo := repositories.NewUserRepositoryMock()
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "dsfwdhuasjkdasdsakjdhak"
		userRepo.On("FindByEmail").Return(nil, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.LoginUser(loginDto)
		// assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, err.(*fiber.Error))
		assert.Equal(t, 401, err.(*fiber.Error).Code)
	})

	t.Run("password incorrect", func(t *testing.T) {
		// mock
		loginDto := models.LoginUserDto{
			Email:    "boat@gmail.com",
			Password: "1234",
		}
		userRepo := repositories.NewUserRepositoryMock()
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "dsfwdhuasjkdasdsakjdhak"
		userRepo.On("FindByEmail").Return(&userCreatedMock1, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, err := userService.LoginUser(loginDto)
		// assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, err.(*fiber.Error))
		assert.Equal(t, 401, err.(*fiber.Error).Code)
	})

	t.Run("login success", func(t *testing.T) {
		// mock
		loginDto := models.LoginUserDto{
			Email:    "boat@gmail.com",
			Password: "1234",
		}
		userRepo := repositories.NewUserRepositoryMock()
		userCreatedMock1.ID = primitive.NewObjectID()
		userCreatedMock1.Password = "1234"
		hash, _ := utils.Hash(userCreatedMock1.Password)
		userCreatedMock1.Password = hash
		userRepo.On("FindByEmail").Return(&userCreatedMock1, nil)
		// test
		userService := services.NewUserService(userRepo)
		result, _ := userService.LoginUser(loginDto)
		// assert
		assert.NotEmpty(t, result)
		assert.IsType(t, &models.Token{}, result)
	})
}