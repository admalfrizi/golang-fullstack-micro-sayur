package entity

type UserEntity struct {
	ID		int64
	Name 	string
	Email 	string
	Password string
	Role 	 string
	Address	 string
	Lat		 string
	Lng		 string
	Phone	 string
	IsVerified bool
}