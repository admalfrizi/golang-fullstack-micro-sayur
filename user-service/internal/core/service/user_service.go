package service

import (
	"context"
	"errors"
	"user-service/config"
	"user-service/internal/adapter/repository"
	"user-service/internal/core/domain/entity"
	"user-service/utils/conv"

	"github.com/labstack/gommon/log"
)

type UserServiceInterface interface {
	SignIn(ctx context.Context, req entity.UserEntity) (*entity.UserEntity, string, error)
}

type userService struct {
	repo repository.UserRepositoryInterface
	cfg *config.Config
	jwtService JwtServiceInterface
}

// SignIn implements UserServiceInterface.
func (u *userService) SignIn(ctx context.Context, req entity.UserEntity) (*entity.UserEntity, string, error) {
	user, err := u.repo.GetUserByEmail(ctx, req.Email) 

	if err != nil {
		log.Errorf("[UserService-1] SignIn : %v", err)
		return nil, "",err
	}

	if checkPass := conv.CheckPasswordHash(req.Password, user.Password); !checkPass {
		err = errors.New("password is incorrect")
		log.Errorf("[UserService-1] SignIn: %v", err)
		return nil, "",err
	}

	token, err := u.jwtService.GenerateToken(user.ID)
	if err != nil {
		log.Errorf("[UserService-1] SignIn: %v", err)
		return nil,"",err
	}

	return user, token, nil
}

func NewUserService(repo repository.UserRepositoryInterface, cfg *config.Config, jwtService JwtServiceInterface) UserServiceInterface {
	return &userService{
		repo: repo,
		cfg: cfg,
		jwtService: jwtService,
	}
}
