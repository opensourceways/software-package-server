package repository

import (
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

type OptToFindSoftwarePkgs struct {
	Status dp.PackageStatus

	PageNum      int
	CountPerPage int
}

type SoftwarePkgIssue struct {
	domain.SoftwarePkgBasicInfo

	domain.SoftwarePkgIssueInfo
}

type ImportedSoftwarePkg struct {
	domain.SoftwarePkgBasicInfo

	domain.ImportedSoftwarePkgInfo
}

type SoftwarePkg interface {
	// AddSoftwarePkg adds a new pkg
	AddSoftwarePkg(*domain.SoftwarePkg) error

	// FindSoftwarePkgIssue find an issue belonging to a pkg
	FindSoftwarePkgIssue(pid string) (SoftwarePkgIssue, error)

	// FindSoftwarePkg find an imported pkg by id
	FindImportedSoftwarePkg(pid string) (ImportedSoftwarePkg, error)

	FindSoftwarePkgs(OptToFindSoftwarePkgs) (r []domain.SoftwarePkgBasicInfo, total int, err error)
}
