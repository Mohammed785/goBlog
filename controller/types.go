package controller

type userCredentials struct {
	Username string `json:"username" form:"username" validate:"required,min=8,max=30"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=30"`
}

