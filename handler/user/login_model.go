package user

type UserLoginRequest struct {
	Mobile      string `json:"Mobile"`
	Description string `json:"description"`
	NickName    string `json:"nickname"`
	Password    string `json:"password"`
}
