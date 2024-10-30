package account

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ResLogin struct {
	Token string `json:"token"`
}
