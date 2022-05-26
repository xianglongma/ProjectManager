package project

type CreateRequest struct {
	Title          string `json:"title,omitempty"`
	Description    string `json:"description,omitempty"`
	Detail         string `json:"detail,omitempty"`
	Url            string `json:"url,omitempty"`
	AvatarUrl      string `json:"avatar_url,omitempty"`
	Permmit        int    `json:"permmit,omitempty"` // 1 全体 2 项目组 3 个人
	Users          string `json:"users,omitempty"`
	Start          int    `json:"start"`
	End            int    `json:"end"`
	AllProcessNode string `json:"all_process_node"`
}
