package repositoryimpl

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	repositoryerr "github.com/opensourceways/software-package-server/common/domain/repository"
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/repository"
)

type softwareImpl struct {
	db *gorm.DB
}

func (s softwareImpl) AddSoftwarePkg(pkg *domain.SoftwarePkg) error {
	if s.exists(pkg.Application.PackageName.PackageName()) {
		return repositoryerr.NewErrorDuplicateCreating(errors.New("package name exists"))
	}

	softwareDO, softIssueDO := s.entityToDO(pkg)
	// TODO issue_num  craete issue
	softwareDO.IssueNum = ""
	softIssueDO.IssueId = 0
	softIssueDO.IssueStatus = ""
	softIssueDO.IssueUrl = ""
	softIssueDO.IssueNum = ""

	return s.save(softwareDO, softIssueDO)
}

func NewSoftware(db *gorm.DB) repository.SoftwarePkg {
	return softwareImpl{db: db}
}

func (s softwareImpl) exists(packageName string) bool {
	var soft SoftwarePkgDO
	err := s.db.Model(soft).Where("package_name = ?", packageName).First(&soft).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (s softwareImpl) save(soft *SoftwarePkgDO, softIssue *SoftwareIssueDO) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(soft).Create(soft).Error
		if err != nil {
			return err
		}

		err = tx.Model(softIssue).Create(softIssue).Error

		return err
	})
}

func (s softwareImpl) entityToDO(pkg *domain.SoftwarePkg) (*SoftwarePkgDO, *SoftwareIssueDO) {
	u := uuid.New()
	t := time.Now()
	softwareDO := &SoftwarePkgDO{
		UUID:            u,
		SourceCode:      pkg.Application.SourceCode.Address.URL(),
		PackageName:     pkg.Application.PackageName.PackageName(),
		PackageDesc:     pkg.Application.PackageDesc.PackageDesc(),
		PackageLicense:  pkg.Application.SourceCode.License.License(),
		PackagePlatform: pkg.Application.PackagePlatform.PackagePlatform(),
		PackageSig:      pkg.Application.ImportingPkgSig.ImportingPkgSig(),
		PackageReason:   pkg.Application.ReasonToImportPkg.ReasonToImportPkg(),
		Status:          pkg.Status.PackageStatus(),
		ApplyUser:       pkg.Importer.Account(),
		CreateTime:      t,
		UpdateTime:      t,
	}

	softIssueDO := &SoftwareIssueDO{
		UUID:          u,
		IssuePlatform: "gitee",
		CreateTime:    t,
		UpdateTime:    t,
	}

	return softwareDO, softIssueDO
}

func (s softwareImpl) FindSoftwarePkgIssue(pid string) (repository.SoftwarePkgIssue, error) {
	//TODO implement me
	panic("implement me")
}

func (s softwareImpl) FindSoftwarePkgs(pkgs repository.OptToFindSoftwarePkgs) (r []domain.SoftwarePkgBasicInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}
