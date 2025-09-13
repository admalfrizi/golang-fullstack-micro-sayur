package request

type SignInRequest struct {
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=8,required"`
}

type SignUpRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=8,required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"min=8,required"`

}