package auth

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,max=320"`
	Name     string `json:"name" binding:"required,max=60"`
	Password string `json:"password" binding:"required,max=72,min=8"`
	Address  string `json:"address"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=320"`
	Password string `json:"password" binding:"required,max=72,min=8"`
}

type LoginResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
