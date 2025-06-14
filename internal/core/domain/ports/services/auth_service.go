package services

import (
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
)

type AuthService interface {
	Register(req entities.RegisterRequest) (*entities.User, error)
	Login(req entities.LoginRequest) (*entities.LoginResponse, error)
	GetUserByID(id uint) (*entities.User, error)
	UpdateUser(user *entities.User) error
}
