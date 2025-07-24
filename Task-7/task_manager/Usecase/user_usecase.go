package Usecase

import (
	"context"
	"task_manager/Domain"
	"task_manager/Infrastructure"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user Domain.User) (Domain.User, error)
	LogIn(ctx context.Context, username, password string) (string, error)
}

type userUsecase struct {
	userRepo        Domain.UserRepository
	jwtService      Infrastructure.JWTService
	passwordService Infrastructure.PasswordService
}

// LogIn implements UserUsecase.
func (u *userUsecase) LogIn(ctx context.Context, username string, password string) (string, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, username)
	if err != nil{
		return "", err
	}

	err = u.passwordService.ComparePassword(user.Password, password)
	if err != nil{
		return "", err
	}
	return u.jwtService.GenerateToken(user.ID.Hex(), user.Username, string(user.Role))
}

// RegisterUser implements UserUsecase.
func (u *userUsecase) RegisterUser(ctx context.Context, user Domain.User) (Domain.User, error) {
	if err := user.Validate(); err != nil{
		return Domain.User{}, err
	}

	hashedPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return Domain.User{}, err
	}

	user.Password = hashedPassword
	return u.userRepo.CreateUser(ctx, user)
}

func NewUserUsecase(userRepo Domain.UserRepository, jwtService Infrastructure.JWTService, passwordService Infrastructure.PasswordService) UserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		jwtService:      jwtService,
		passwordService: passwordService,
	}
}
