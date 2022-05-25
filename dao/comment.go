package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ParentID     int    `json:"parent_id"`
	ReplierID    int    `json:"replier_id"`
	ResourceID   int    `json:"resource_id"`   // 被评论的资源id
	ResourceType int    `json:"resource_type"` // 被评论的资源type
	Content      string `json:"content"`       // 评论内容
	Permmit      int    `json:"permmit"`       // 是否被禁用
}
