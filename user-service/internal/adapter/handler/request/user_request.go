package request

type SignInRequest struct {
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=8,required"`
}