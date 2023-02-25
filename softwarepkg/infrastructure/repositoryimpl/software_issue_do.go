package repositoryimpl

import (
	"time"

	"github.com/google/uuid"
)

type SoftwareIssueDO struct {
	UUID          uuid.UUID `json:"uuid" gorm:"column:uuid;type:uuid"`
	IssueUrl      string    `json:"issue_url" gorm:"column:issue_url"`
	IssueNum      string    `json:"issue_num" gorm:"column:issue_num"`
	IssuePlatform string    `json:"issue_platform" gorm:"column:issue_platform"`
	IssueStatus   string    `json:"issue_status" gorm:"column:issue_status"`
	IssueId       int       `json:"issue_id" gorm:"column:issue_id"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime    time.Time `json:"update_time" gorm:"column:update_time"`
}

func (SoftwareIssueDO) TableName() string {
	return "software_issue"
}
