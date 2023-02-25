package repositoryimpl

import (
	"time"

	"github.com/google/uuid"
)

type SoftwareIssueCommentDO struct {
	UUID       uuid.UUID `json:"uuid" gorm:"column:uuid;type:uuid"`
	Content    string    `json:"content" gorm:"column:content"`
	ApplyUser  string    `json:"apply_user" gorm:"column:apply_user"`
	Status     int       `json:"status" gorm:"column:status"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

func (SoftwareIssueCommentDO) TableName() string {
	return "software_issue_comment"
}
