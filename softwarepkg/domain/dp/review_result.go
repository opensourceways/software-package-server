package dp

import "errors"

const (
	pkgReviewResultRejected = "rejected"
	pkgReviewResultApproved = "approved"
)

var (
	PkgReviewResultRejected = packageReviewResult(pkgReviewResultRejected)
	PkgReviewResultApproved = packageReviewResult(pkgReviewResultApproved)
)

type PackageReviewResult interface {
	PackageReviewResult() string
}

type packageReviewResult string

func NewPackageReviewResult(v string) (PackageReviewResult, error) {
	if v == "" {
		return nil, errors.New("empty review result")
	}

	return packageReviewResult(v), nil
}

func (v packageReviewResult) PackageReviewResult() string {
	return string(v)
}
