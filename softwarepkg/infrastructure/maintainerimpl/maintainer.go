package maintainerimpl

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/maintainer"
)

func NewMaintainerImpl(cfg *Config) maintainer.Maintainer {
	return maintainerImpl{
		permissionURL: cfg.PermissionURL,
		community:     cfg.Community,
		cli:           utils.NewHttpClient(3),
	}
}

type maintainerImpl struct {
	permissionURL string
	community     string
	cli           utils.HttpClient
}

type SigPermission struct {
	Data []struct {
		Sig  string   `json:"sig"`
		Type []string `json:"type"`
	} `json:"data"`
}

func (s SigPermission) checkPermission(sig string) (flag bool) {
	for _, v := range s.Data {
		if v.Sig == sig && strings.Contains(strings.Join(v.Type, ","), "maintainers") {
			flag = true
			break
		}
	}

	return
}

func (m maintainerImpl) baseUrl(user string) string {
	return fmt.Sprintf(m.permissionURL, m.community, user)
}

func (m maintainerImpl) HasPermission(info *domain.SoftwarePkgBasicInfo, user dp.Account) (
	flag bool, err error,
) {
	req, err := http.NewRequest("GET", m.baseUrl(user.Account()), nil)
	if err != nil {
		return
	}

	var res SigPermission
	_, err = m.cli.ForwardTo(req, &res)
	if err != nil {
		return
	}

	sig := info.Application.ImportingPkgSig.ImportingPkgSig()

	flag = res.checkPermission(sig)

	return
}
