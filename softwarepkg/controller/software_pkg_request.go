package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/opensourceways/software-package-server/common/middleware"
	"github.com/opensourceways/software-package-server/softwarepkg/app"
	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

const (
	pageNum      = 1
	countPerPage = 10
)

type softwarePkgRequest struct {
	SourceCodeUrl     string `json:"source_code"     binding:"required"`
	SourceCodeLicense string `json:"license"         binding:"required"`
	PackageName       string `json:"pkg_name"        binding:"required"`
	PackageDesc       string `json:"desc"            binding:"required"`
	PackagePlatform   string `json:"platform"        binding:"required"`
	PackageSig        string `json:"sig"             binding:"required"`
	PackageReason     string `json:"reason"          binding:"required"`
}

func (s softwarePkgRequest) toCmd(importer *domain.User) (
	cmd app.CmdToApplyNewSoftwarePkg, err error,
) {
	cmd.Importer = *importer

	cmd.PkgName, err = dp.NewPackageName(s.PackageName)
	if err != nil {
		return
	}

	application := &cmd.Application

	application.SourceCode.Address, err = dp.NewURL(s.SourceCodeUrl)
	if err != nil {
		return
	}

	application.SourceCode.License, err = dp.NewLicense(s.SourceCodeLicense)
	if err != nil {
		return
	}

	application.ImportingPkgSig, err = dp.NewImportingPkgSig(s.PackageSig)
	if err != nil {
		return
	}

	application.ReasonToImportPkg, err = dp.NewReasonToImportPkg(s.PackageReason)
	if err != nil {
		return
	}

	application.PackageDesc, err = dp.NewPackageDesc(s.PackageDesc)
	if err != nil {
		return
	}

	application.PackagePlatform, err = dp.NewPackagePlatform(s.PackagePlatform)

	return
}

type softwarePkgListQuery struct {
	Importer     string `json:"importer"       form:"importer"`
	Phase        string `json:"phase"          form:"phase"`
	PageNum      int    `json:"page_num"       form:"page_num"`
	CountPerPage int    `json:"count_per_page" form:"count_per_page"`
}

func (s softwarePkgListQuery) toCmd() (pkg app.CmdToListPkgs, err error) {
	if s.Importer != "" {
		if pkg.Importer, err = dp.NewAccount(s.Importer); err != nil {
			return
		}
	}

	if s.Phase != "" {
		if pkg.Phase, err = dp.NewPackagePhase(s.Phase); err != nil {
			return
		}
	}

	if s.PageNum > 0 {
		pkg.PageNum = s.PageNum
	} else {
		pkg.PageNum = pageNum
	}

	if s.CountPerPage > 0 {
		pkg.CountPerPage = s.CountPerPage
	} else {
		pkg.CountPerPage = countPerPage
	}

	return
}

func toDomainUser(ctx *gin.Context) (user domain.User, err error) {
	if user.Email, err = dp.NewEmail(middleware.GetEmail(ctx)); err != nil {
		return
	}

	var u string
	if u = middleware.GiteeUserName(ctx); len(u) == 0 {
		u = middleware.UserName(ctx)
	}
	user.Account, err = dp.NewAccount(u)

	return
}
