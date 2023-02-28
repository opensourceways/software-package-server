package repositoryimpl

import (
	"errors"

	"gorm.io/gorm"

	commonrepo "github.com/opensourceways/software-package-server/common/domain/repository"
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/repository"
)

type softwarePkgImpl struct {
	db *gorm.DB
}

func NewSoftwarePkg(db *gorm.DB) repository.SoftwarePkg {
	return softwarePkgImpl{db: db}
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
	softwareDO := s.toSoftwareDO(pkg)

	return s.save(softwareDO)
}

func (s softwarePkgImpl) save(soft *SoftwarePkgDO) error {
	query := s.db.Where(&SoftwarePkgDO{PackageName: soft.PackageName}).FirstOrCreate(soft)
	if err := query.Error; err != nil {
		return err
	}

	if query.RowsAffected == 0 {
		return commonrepo.NewErrorDuplicateCreating(errors.New("package exists"))
	}

	return nil
}
