package translationimpl

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

func NewTranslationService(cfg *Config) service {
	return service{
		cli: utils.NewHttpClient(3),
		cfg: *cfg,
	}
}

type translation struct {
	SrcText        string `json:"src_text"`
	TranslatedText string `json:"translated_text"`
	From           string `json:"from"`
	To             string `json:"to"`
}

type service struct {
	cli utils.HttpClient
	cfg Config
}

func (s service) token() (string, error) {
	return gentoken(&s.cfg.Cloud)
}

func gentoken(cfg *CloudConfig) (string, error) {
	str := `
{ 
    "auth": { 
        "identity": { 
            "methods": [ 
                "password" 
            ], 
            "password": { 
                "user": { 
                    "name": "%s", 
                    "password": "%s", 
                    "domain": { 
                        "name": "%s" 
                    } 
                } 
            } 
        }, 
        "scope": { 
            "project": { 
                "name": "%s" 
            } 
        } 
    } 
}
`
	body := fmt.Sprintf(
		str, cfg.User, cfg.Password, cfg.Domain, cfg.Project,
	)

	resp, err := http.Post(
		cfg.AuthEndpoint, "application/json",
		strings.NewReader(body),
	)
	if err != nil {
		return "", err
	}

	t := resp.Header.Get("X-Subject-Token")

	resp.Body.Close()

	return t, nil
}

func (s service) Translation(content dp.ReviewComment, from, to string) (text string, err error) {
	str := `
{
    "text":"%s",
    "from":"%s",
    "to":"%s"
}
`
	body := fmt.Sprintf(str, content.ReviewComment(), from, to)

	token, err := s.token()
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		http.MethodPost, s.cfg.TranslationURL, strings.NewReader(body),
	)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", token)

	var res translation
	if _, err = s.cli.ForwardTo(req, &res); err != nil {
		return
	}

	return res.TranslatedText, nil
}
