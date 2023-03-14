package translation

import "github.com/opensourceways/software-package-server/softwarepkg/domain/dp"

type Translating interface {
	Translate(string, dp.Language) (string, error)
}
