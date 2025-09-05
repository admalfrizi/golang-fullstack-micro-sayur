package repository

import (
	"context"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/domain/model"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type userRepository struct {
	db *gorm.DB
}

// GetUserByEmail implements UserRepositoryInterface.
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	modelUser := model.User{}

	if err := u.db.Where("email = ? && is_verified = ?", email, true).Preload("Roles").First(&modelUser).Error; err != nil {
		log.Errorf("[UserRepository-1] GetUserByEmail: %v", err)
		return nil,err
	}

	return &entity.UserEntity{
		ID: modelUser.ID,
		Name: modelUser.Name,
		Email: modelUser.Email,
		Password: modelUser.Password,
		Role: modelUser.Roles[0].Name,
		Address: modelUser.Address,
		Lat: modelUser.Lat,
		Lng: modelUser.Lng,
		Phone: modelUser.Phone,
		IsVerified: modelUser.IsVerified,
	},nil
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: db}
}
