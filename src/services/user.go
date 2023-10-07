package services

import (
	"ecommerce-backend/src/configs"
	"ecommerce-backend/src/logs"
	"ecommerce-backend/src/models"
	"ecommerce-backend/src/repositories"
	"ecommerce-backend/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(req models.User) (*models.User, error)
	LoginUser(req models.LoginUserDto) (*models.Token, error)
	GetProfile(_id string) (*models.User, error)
}
type userService struct {
	userRepo repositories.UserRopository
}

func NewUserService(userRepo repositories.UserRopository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(req models.User) (*models.User, error) {
	// check exist
	currentUser, err := s.userRepo.FindByEmail(req.Email, false)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if currentUser != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "email already exist 😜")
	}
	// hash password
	hash, err := utils.Hash(req.Password)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	req.Password = hash

	user := models.NewUser(req)
	newUser, err := s.userRepo.Create(user)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return newUser, nil
}

func (s userService) LoginUser(req models.LoginUserDto) (*models.Token, error) {
	currentUser, err := s.userRepo.FindByEmail(req.Email, true)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if currentUser == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Email or Password are not correct 🥲")
	}
	// check password
	ok := utils.CompareHash(req.Password, currentUser.Password)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Email or Password are not correct 🥲")
	}

	claims := jwt.MapClaims{
		"id":    currentUser.ID,
		"email": currentUser.Email,
		"role":  currentUser.Role,
	}
	token, err := utils.Token(claims, configs.Cfg.Jwt.Secret)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return &token, nil
}

func (s userService) GetProfile(id string) (*models.User, error) {
	// convert id string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid object id")
	}

	currentUser, err := s.userRepo.FindByID(objectId, false)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if currentUser == nil {
		return nil, fiber.ErrNotFound
	}

	return currentUser, nil
}
