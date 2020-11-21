package user

type RegisterUserInput struct { //Untuk Mewakkili apa yang diinputkan oleh user
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email", binding:"required, email"`
	Password string `json:"password", binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email", binding:"required, email"`
}
