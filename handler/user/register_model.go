package user

type RegisterRequest struct {
	Mobile      string `json:"Mobile"`
	Description string `json:"description"`
	NickName    string `json:"nickname"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}
