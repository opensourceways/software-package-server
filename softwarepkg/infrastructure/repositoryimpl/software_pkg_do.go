package repositoryimpl

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type SoftwarePkgDO struct {
	UUID            uuid.UUID      `json:"uuid" gorm:"column:uuid;type:uuid"`
	SourceCode      string         `json:"source_code" gorm:"column:source_code"`
	PackageName     string         `json:"package_name" gorm:"column:package_name"`
	PackageDesc     string         `json:"package_desc" gorm:"column:package_desc"`
	PackageLicense  string         `json:"package_license" gorm:"column:package_license"`
	PackagePlatform string         `json:"package_platform" gorm:"column:package_platform"`
	PackageSig      string         `json:"package_sig" gorm:"column:package_sig"`
	PackageReason   string         `json:"package_reason" gorm:"column:package_reason"`
	Status          string         `json:"status" gorm:"column:status"`
	ApplyUser       string         `json:"apply_user" gorm:"column:apply_user"`
	IssueNum        string         `json:"issue_num" gorm:"column:issue_num"`
	PackageRepoLink string         `json:"package_repo_link" gorm:"column:package_repo_link"`
	ApproveUser     pq.StringArray `json:"approve_user" gorm:"column:approve_user;type:text[];default:'{}'"`
	RejectUser      pq.StringArray `json:"reject_user" gorm:"column:reject_user;type:text[];default:'{}'"`
	CreateTime      time.Time      `json:"create_time" gorm:"column:create_time"`
	UpdateTime      time.Time      `json:"update_time" gorm:"column:update_time"`
}

func (SoftwarePkgDO) TableName() string {
	return "software_pkg"
}
