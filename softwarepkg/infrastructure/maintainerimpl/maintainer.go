package maintainerimpl

import (
	"fmt"
	"net/http"

	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/maintainer"
)

func NewMmaintainerImpl(cfg *Config) maintainer.Maintainer {
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

func (m maintainerImpl) baseUrl(user string) string {
	return fmt.Sprintf(m.permissionURL, m.community, user)
}

func (m maintainerImpl) HasPermission(info *domain.SoftwarePkgBasicInfo, user dp.Account) (flag bool, err error) {
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
	for _, v := range res.Data {
		if v.Sig == sig {
			for _, s := range v.Type {
				if s == "maintainers" {
					flag = true
					return
				}
			}
		}
	}

	return
}
