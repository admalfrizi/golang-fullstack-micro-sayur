package entity

import "time"

type VerificationTokenEntity struct {
	ID        	int64
	UserID    	int64
	Token     	string
	TokenType 	string
	ExpiresAt 	time.Time
	User		UserEntity
}