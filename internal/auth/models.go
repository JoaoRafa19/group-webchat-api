package auth 


type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

