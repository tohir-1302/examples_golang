package request

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Id       int    `json:"id" validate:"required,gte=1"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
