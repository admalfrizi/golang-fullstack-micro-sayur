package model

import "time"

type User struct {
	ID         int64 `gorm:primaryKey`
	Name       string
	Email      string
	Password   string
	Address    string
	Phone      string
	Photo      string
	Lat        string
	Lng        string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt	time.Time
	DeletedAt	time.Time
	Roles		[]Role `gorm:"many2many:user_roles"`
}