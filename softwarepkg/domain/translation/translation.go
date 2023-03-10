package translation

import "github.com/opensourceways/software-package-server/softwarepkg/domain/dp"

type Translation interface {
	Translation(content dp.ReviewComment, from, to string) (text string, err error)
}
