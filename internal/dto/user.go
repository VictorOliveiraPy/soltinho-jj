package dto

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	RoleID   string `json:"role_id"`
	Active   bool   `json:"active"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

type GetJWTOutput struct {
	AccessToken string `json:"acess_token"`
}
