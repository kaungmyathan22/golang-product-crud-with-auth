package dto

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3" error:"username field is required and must be a minimum of 3."`
	Password string `json:"password" validate:"required,min=6" error:"password field is required and must be a minimum of 6."`
}

type UpdateUserDTO struct {
	Username string `json:"username"`
}

type UpdatePasswordDTO struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
