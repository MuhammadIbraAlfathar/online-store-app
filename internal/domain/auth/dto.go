package auth

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,max=320"`
	Name     string `json:"name" binding:"required,max=60"`
	Password string `json:"password" binding:"required,max=72,min=8"`
	Address  string `json:"address"`
}
