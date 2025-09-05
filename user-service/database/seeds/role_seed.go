package seeds

import (
	"log"
	"user-service/internal/core/domain/model"

	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []model.Role{
		{
			Name: "Super Admin",
		},
		{
			Name: "Customer",
		},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, model.Role{Name: role.Name}).Error; err != nil {
			log.Fatalf("%s : %v", err.Error(), err)
		} else {
			log.Printf("Role %s created", role.Name)
		}
	}
}