package controller

import (
	"github.com/opensourceways/software-package-server/softwarepkg/app"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

type softwareRequest struct {
	Username          string `json:"username" binding:"required"`
	SourceCodeUrl     string `json:"source_code_url" binding:"required"`
	SourceCodeLicense string `json:"source_code_license" binding:"required"`
	PackageName       string `json:"package_name" binding:"required"`
	PackageDesc       string `json:"package_desc" binding:"required"`
	PackagePlatform   string `json:"package_platform" binding:"required"`
	PackageSig        string `json:"package_sig" binding:"required"`
	PackageReason     string `json:"package_reason" binding:"required"`
}

func (s softwareRequest) ToCmd() (pkg app.CmdToApplyNewSoftwarePkg, user dp.Account, err error) {
	pkg.SourceCode.Address, err = dp.NewURL(s.SourceCodeUrl)
	if err != nil {
		return
	}

	pkg.SourceCode.License, err = dp.NewLicense(s.SourceCodeLicense)
	if err != nil {
		return
	}

	pkg.ImportingPkgSig, err = dp.NewImportingPkgSig(s.PackageSig)
	if err != nil {
		return
	}

	pkg.ReasonToImportPkg, err = dp.NewReasonToImportPkg(s.PackageReason)
	if err != nil {
		return
	}

	pkg.PackageName, err = dp.NewPackageName(s.PackageName)
	if err != nil {
		return
	}

	pkg.PackageDesc, err = dp.NewPackageDesc(s.PackageDesc)
	if err != nil {
		return
	}

	user, err = dp.NewAccount(s.Username)
	if err != nil {
		return
	}

	pkg.PackagePlatform, err = dp.NewPackagePlatform(s.PackagePlatform)

	return
}
