package user

type UserUpdateInfoRequest struct {
	NickName    string `json:"nick_name,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email,omitempty"`
	Mobile      string `json:"mobile,omitempty"`
}
