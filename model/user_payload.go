package model

type UserLoginPayload struct {
	Email    string   `json:"email" form:"email"`
	Password Password `json:"pass" form:"pass"`
}

type UserRegisterPayload struct {
	FirstName string   `json:"first_name" form:"first_name"`
	LastName  string   `json:"last_name" form:"last_name"`
	Email     string   `json:"email" form:"email" validate:"required,email"`
	Password  Password `json:"pass" form:"pass"`
}

type UserUpdatePayload struct {
	FirstName *string `json:"first_name" form:"first_name"`
	LastName  *string `json:"last_name" form:"last_name"`
	Email     *string `json:"email" form:"email" validate:"required,email"`
}
