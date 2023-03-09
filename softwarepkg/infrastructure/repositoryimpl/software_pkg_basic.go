package repositoryimpl

import (
	"github.com/google/uuid"

	commonrepo "github.com/opensourceways/software-package-server/common/domain/repository"
	"github.com/opensourceways/software-package-server/common/infrastructure/postgresql"
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/repository"
)

// softwarePkgBasic
type softwarePkgBasic struct {
	basicDBCli dbClient
}

func (s softwarePkgBasic) SaveSoftwarePkg(pkg *domain.SoftwarePkgBasicInfo, version int) error {
	u, err := uuid.Parse(pkg.Id)
	if err != nil {
		return err
	}

	filter := SoftwarePkgBasicDO{Id: u, Version: version}
	updates := s.toSoftwarePkgUpdate(pkg)

	if err = s.basicDBCli.UpdateRecord(&filter, updates); err != nil && s.basicDBCli.IsRowNotFound(err) {
		err = commonrepo.NewErrorResourceNotFound(err)
	}

	return err
}

func (s softwarePkgBasic) FindSoftwarePkgBasicInfo(pid string) (
	info domain.SoftwarePkgBasicInfo, version int, err error,
) {
	v, err := uuid.Parse(pid)
	if err != nil {
		return
	}

	var do SoftwarePkgBasicDO

	if err = s.basicDBCli.GetRecord(&SoftwarePkgBasicDO{Id: v}, &do); err != nil {
		if s.basicDBCli.IsRowNotFound(err) {
			err = commonrepo.NewErrorResourceNotFound(err)
		}
	} else {
		version = do.Version

		info, err = do.toSoftwarePkgBasicInfo()
	}

	return
}

func (s softwarePkgBasic) FindSoftwarePkgs(pkgs repository.OptToFindSoftwarePkgs) (
	r []domain.SoftwarePkgBasicInfo, total int, err error,
) {
	var filter SoftwarePkgBasicDO
	if pkgs.Importer != nil {
		filter.Importer = pkgs.Importer.Account()
	}

	if pkgs.Phase != nil {
		filter.Phase = pkgs.Phase.PackagePhase()
	}

	if total, err = s.basicDBCli.Count(&filter); err != nil || total == 0 {
		return
	}

	var dos []SoftwarePkgBasicDO

	err = s.basicDBCli.GetRecords(
		&filter, &dos,
		postgresql.Pagination{
			PageNum:      pkgs.PageNum,
			CountPerPage: pkgs.CountPerPage,
		},
		[]postgresql.SortByColumn{
			{Column: fieldAppliedAt},
		},
	)
	if err != nil || len(dos) == 0 {
		return
	}

	r = make([]domain.SoftwarePkgBasicInfo, len(dos))
	for i := range dos {
		if r[i], err = dos[i].toSoftwarePkgBasicInfo(); err != nil {
			return
		}
	}

	return
}

func (s softwarePkgBasic) AddSoftwarePkg(pkg *domain.SoftwarePkgBasicInfo) error {
	var do SoftwarePkgBasicDO
	s.toSoftwarePkgBasicDO(pkg, &do)

	pkg.Id = do.Id.String()

	err := s.basicDBCli.Insert(
		&SoftwarePkgBasicDO{
			PackageName: do.PackageName,
		},
		&do,
	)
	if err != nil && s.basicDBCli.IsRowExists(err) {
		err = commonrepo.NewErrorDuplicateCreating(err)
	}

	return err
}
