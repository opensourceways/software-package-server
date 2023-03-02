package repositoryimpl

import (
	"sort"

	"github.com/google/uuid"

	commonrepo "github.com/opensourceways/software-package-server/common/domain/repository"
	"github.com/opensourceways/software-package-server/common/infrastructure/postgresql"
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/repository"
)

type softwarePkgImpl struct {
	cli dbClient
}

func NewSoftwarePkg(cli dbClient) repository.SoftwarePkg {
	return softwarePkgImpl{cli: cli}
}

func (s softwarePkgImpl) SaveSoftwarePkg(pkg *domain.SoftwarePkgBasicInfo, version int) error {
	//TODO implement me
	return nil
}

func (s softwarePkgImpl) FindSoftwarePkgBasicInfo(pid string) (domain.SoftwarePkgBasicInfo, int, error) {
	//TODO implement me
	return domain.SoftwarePkgBasicInfo{}, 0, nil
}

func (s softwarePkgImpl) FindSoftwarePkg(pid string) (
	pkg domain.SoftwarePkg, version int, err error,
) {
	var u uuid.UUID
	if u, err = uuid.Parse(pid); err != nil {
		return
	}

	var (
		result SoftwarePkgPreload
		filter = SoftwarePkgDO{UUID: u}
	)
	if err = s.cli.GetTableRecord(&filter, &result, "SoftwarePkgReviewDO"); err != nil {
		return
	}

	version = result.SoftwarePkgDO.Version

	if pkg.SoftwarePkgBasicInfo, err = result.SoftwarePkgDO.toSoftwarePkgSummary(); err != nil {
		return
	}

	pkg.Comments = make([]domain.SoftwarePkgReviewComment, len(result.SoftwarePkgReviewDO))

	for i, do := range result.SoftwarePkgReviewDO {
		if pkg.Comments[i], err = do.toSoftwarePkgReviewCommentSummary(); err != nil {
			return
		}
	}

	sort.Slice(pkg.Comments, func(i, j int) bool {
		return pkg.Comments[i].CreatedAt < pkg.Comments[j].CreatedAt
	})

	return
}

func (s softwarePkgImpl) FindSoftwarePkgs(pkgs repository.OptToFindSoftwarePkgs) (
	r []domain.SoftwarePkgBasicInfo, total int, err error,
) {
	var filter SoftwarePkgDO
	if pkgs.Importer != nil {
		filter.ImportUser = pkgs.Importer.Account()
	}

	if pkgs.Phase != nil {
		filter.Phase = pkgs.Phase.PackagePhase()
	}

	if total, err = s.cli.Counts(&filter); err != nil || total == 0 {
		return
	}

	var sort = []postgresql.SortByColumn{
		{Column: applyTime, Ascend: false},
	}

	var p = postgresql.Pagination{PageNum: pkgs.PageNum, CountPerPage: pkgs.CountPerPage}

	var result []SoftwarePkgDO
	if err = s.cli.GetTableRecords(&filter, &result, p, sort); err != nil {
		return
	}

	r = make([]domain.SoftwarePkgBasicInfo, len(result))
	for i, v := range result {
		if r[i], err = v.toSoftwarePkgSummary(); err != nil {
			return
		}
	}

	return
}

func (s softwarePkgImpl) AddReviewComment(pid string, comment *domain.SoftwarePkgReviewComment) error {
	//TODO implement me
	return nil
}

func (s softwarePkgImpl) AddSoftwarePkg(pkg *domain.SoftwarePkgBasicInfo) error {
	v := s.toSoftwarePkgDO(pkg)
	filter := SoftwarePkgDO{PackageName: pkg.PkgName.PackageName()}
	err := s.cli.Insert(&filter, &v)
	if err != nil && s.cli.IsRowExists(err) {
		err = commonrepo.NewErrorDuplicateCreating(err)
	}

	return err
}
