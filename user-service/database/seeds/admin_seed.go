package seeds

import (
	"log"
	"user-service/internal/core/domain/model"
	"user-service/utils/conv"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	modelRole := model.Role{}
	err = db.Where("name = ?", "Super Admin").First(&modelRole).Error

	if err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	}

	admin := model.User{
		Name: "super admin",
		Email: "superadmin@gmail.com",
		Password: bytes,
		IsVerified: true,
		Roles: []model.Role{modelRole},
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "superadmin@gmail.com"}).Error; err != nil {
		log.Fatalf("%s: %v", err.Error(), err)
	} else {
		log.Printf("Admin %s created", admin.Name)
	}

	
}