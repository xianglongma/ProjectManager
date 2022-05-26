package comment

type CreateRequest struct {
	Content      string `json:"content,omitempty"`
	ResourceID   int    `json:"resource_id,omitempty" query:"resource_id"`
	ResourceType string `json:"resource_type,omitempty" query:"resource_type"`
}

type QueryListRequest struct {
	ResourceID   int    `json:"resource_id,omitempty" query:"resource_id"`
	ResourceType string `json:"resource_type,omitempty" query:"resource_type"`
}

type QueryListResponse struct {
	Items []QueryListResponseItem `json:"items,omitempty"`
}

type QueryListResponseItem struct {
	Avatar      string `json:"avatar,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	Content     string `json:"content,omitempty"`
	CommentTime string `json:"comment_time,omitempty"`
}
