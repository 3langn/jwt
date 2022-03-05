package dto

//RegisterDTO is used when client post a new user
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1" example:"John Doe"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email" example:"3langn@gmail.com"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6" example:"password"`
}
