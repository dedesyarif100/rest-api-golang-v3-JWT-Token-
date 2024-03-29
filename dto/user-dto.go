package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}


// type UserCreateDTO struct {
// 	ID       uint64 `json:"id" form:"id" binding:"required"`
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required,email" validate:"email"` 
// 	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
// }
