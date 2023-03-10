package dp

import (
	"errors"
	"fmt"

	"github.com/opensourceways/software-package-server/utils"
)

type ReasonToImportPkg interface {
	ReasonToImportPkg() string
}

func NewReasonToImportPkg(v string) (ReasonToImportPkg, error) {
	if v == "" {
		return nil, errors.New("empty reason")
	}

	if max := config.MaxLengthOfReasonToImportPkg; utils.StrLen(v) > max {
		return nil, fmt.Errorf(
			"the length of reason should be less than %d", max,
		)
	}

	return reasonToImportPkg(v), nil
}

type reasonToImportPkg string

func (v reasonToImportPkg) ReasonToImportPkg() string {
	return string(v)
}
