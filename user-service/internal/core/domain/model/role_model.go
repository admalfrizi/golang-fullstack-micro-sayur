package model

import "time"

type Role struct {
	ID         int64 `gorm:primaryKey`
	Name       string
	Users 	   []User `gorm:"many2many:user_roles"`
	CreatedAt  time.Time
	UpdatedAt	time.Time
	DeletedAt	time.Time
}