package maintainerimpl

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/software-package-server/softwarepkg/domain"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

type sigPermission struct {
	Data []struct {
		Sig  string   `json:"sig"`
		Type []string `json:"type"`
	} `json:"data"`
}

func (s sigPermission) hasPermission(sig string) bool {
	for _, v := range s.Data {
		if v.Sig == sig && strings.Contains(strings.Join(v.Type, ","), "maintainers") {
			return true
		}
	}

	return false
}

func NewMaintainerImpl(cfg *Config) maintainerImpl {
	return maintainerImpl{
		cfg: *cfg,
		cli: utils.NewHttpClient(3),
	}
}

type maintainerImpl struct {
	cfg Config
	cli utils.HttpClient
}

func (impl maintainerImpl) baseUrl(user string) string {
	return fmt.Sprintf(impl.cfg.PermissionURL, user)
}

func (impl maintainerImpl) HasPermission(info *domain.SoftwarePkgBasicInfo, user dp.Account) (
	bool, error,
) {
	req, err := http.NewRequest(http.MethodGet, impl.baseUrl(user.Account()), nil)
	if err != nil {
		return false, err
	}

	var res sigPermission
	if _, err = impl.cli.ForwardTo(req, &res); err != nil {
		return false, err
	}

	sig := info.Application.ImportingPkgSig.ImportingPkgSig()
	
	return res.hasPermission(sig), nil
}
