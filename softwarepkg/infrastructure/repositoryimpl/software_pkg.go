package repositoryimpl

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	repositoryerr "github.com/opensourceways/software-package-server/common/domain/repository"
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/repository"
	"github.com/opensourceways/software-package-server/utils"
)

const (
	PackageName = "package_name"
)

type softwarePkgImpl struct {
	db *gorm.DB
}

func (s softwarePkgImpl) SaveSoftwarePkg(pkg *domain.SoftwarePkgBasicInfo, version int) error {
	//TODO implement me
	return nil
}

func (s softwarePkgImpl) FindSoftwarePkgBasicInfo(pid string) (domain.SoftwarePkgBasicInfo, int, error) {
	//TODO implement me
	return domain.SoftwarePkgBasicInfo{}, 0, nil
}

func (s softwarePkgImpl) FindSoftwarePkg(pid string) (domain.SoftwarePkg, int, error) {
	//TODO implement me
	return domain.SoftwarePkg{}, 0, nil
}

func (s softwarePkgImpl) FindSoftwarePkgs(pkgs repository.OptToFindSoftwarePkgs) (r []domain.SoftwarePkgBasicInfo, total int, err error) {
	//TODO implement me
	return nil, 0, err
}

func (s softwarePkgImpl) AddReviewComment(pid string, comment *domain.SoftwarePkgReviewComment) error {
	//TODO implement me
	return nil
}

func (s softwarePkgImpl) AddSoftwarePkg(pkg *domain.SoftwarePkgBasicInfo) error {
	if s.exists(pkg.Application.PackageName.PackageName()) {
		return repositoryerr.NewErrorDuplicateCreating(errors.New("package name exists"))
	}

	softwareDO := s.toSoftwareDO(pkg)

	return s.save(softwareDO)
}

func NewSoftwarePkg(db *gorm.DB) repository.SoftwarePkg {
	return softwarePkgImpl{db: db}
}

func (s softwarePkgImpl) exists(packageName string) bool {
	var soft SoftwarePkgDO
	err := s.db.Model(soft).Where(PackageName+" = ?", packageName).First(&soft).Error
	if err != nil {
		return false
	}

	return true
}

func (s softwarePkgImpl) save(soft *SoftwarePkgDO) error {
	return s.db.Model(soft).Create(soft).Error
}

func (s softwarePkgImpl) toSoftwareDO(pkg *domain.SoftwarePkgBasicInfo) *SoftwarePkgDO {
	t := utils.Now()
	softwareDO := &SoftwarePkgDO{
		UUID:            uuid.New(),
		SourceCode:      pkg.Application.SourceCode.Address.URL(),
		PackageName:     pkg.Application.PackageName.PackageName(),
		PackageDesc:     pkg.Application.PackageDesc.PackageDesc(),
		PackageLicense:  pkg.Application.SourceCode.License.License(),
		PackagePlatform: pkg.Application.PackagePlatform.PackagePlatform(),
		PackageSig:      pkg.Application.ImportingPkgSig.ImportingPkgSig(),
		PackageReason:   pkg.Application.ReasonToImportPkg.ReasonToImportPkg(),
		Phase:           pkg.Phase.PackagePhase(),
		// TODO pkg.Importer.Account() is nil
		ImportUser: "",
		CreateTime: t,
		UpdateTime: t,
	}

	return softwareDO
}
