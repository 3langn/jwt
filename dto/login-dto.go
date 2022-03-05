package dto

type LoginDTO struct {
	Email    string `json:"email" example:"3langn@gmail.com" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" example:"password" form:"password" binding:"required" validate:"min:6"`
}
