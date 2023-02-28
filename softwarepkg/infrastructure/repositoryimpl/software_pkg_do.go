package repositoryimpl

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type SoftwarePkgDO struct {
	UUID            uuid.UUID      `gorm:"column:uuid;type:uuid"`
	SourceCode      string         `gorm:"column:source_code"`
	PackageName     string         `gorm:"column:package_name"`
	PackageDesc     string         `gorm:"column:package_desc"`
	PackageLicense  string         `gorm:"column:package_license"`
	PackagePlatform string         `gorm:"column:package_platform"`
	PackageSig      string         `gorm:"column:package_sig"`
	PackageReason   string         `gorm:"column:package_reason"`
	Phase           string         `gorm:"column:phase"`
	ImportUser      string         `gorm:"column:import_user"`
	Version         int            `gorm:"column:version"`
	PackageRepoLink string         `gorm:"column:package_repo_link"`
	ApproveUser     pq.StringArray `gorm:"column:approve_user;type:text[];default:'{}'"`
	RejectUser      pq.StringArray `gorm:"column:reject_user;type:text[];default:'{}'"`
	CreateTime      int64          `gorm:"column:create_time"`
	UpdateTime      int64          `gorm:"column:update_time"`
}

func (SoftwarePkgDO) TableName() string {
	return "software_pkg"
}
