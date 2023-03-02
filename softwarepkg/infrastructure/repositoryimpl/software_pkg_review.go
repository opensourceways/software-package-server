package repositoryimpl

import (
	"github.com/google/uuid"

	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

type SoftwarePkgReviewDO struct {
	UUID         uuid.UUID `gorm:"column:uuid;type:uuid"`
	Content      string    `gorm:"column:content"`
	ApplyUser    string    `gorm:"column:apply_user"`
	SoftwareUUID string    `gorm:"column:software_uuid;type:uuid"`
	Status       int       `gorm:"column:status"`
	Version      int       `gorm:"column:version"`
	CreateTime   int64     `gorm:"column:create_time"`
	UpdateTime   int64     `gorm:"column:update_time"`
}

func (SoftwarePkgReviewDO) TableName() string {
	return "software_pkg_review"
}

func (s SoftwarePkgReviewDO) toSoftwarePkgReviewCommentSummary() (pkgComment domain.SoftwarePkgReviewComment, err error) {
	pkgComment.CreatedAt = s.CreateTime

	pkgComment.Id = s.UUID.String()

	if pkgComment.Author, err = dp.NewAccount(s.ApplyUser); err != nil {
		return
	}

	pkgComment.Content, err = dp.NewReviewComment(s.Content)

	return
}
