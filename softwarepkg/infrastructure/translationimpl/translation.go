package translationimpl

import (
	"errors"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	v2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nlp/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nlp/v2/model"

	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

func NewTranslationService(cfg *Config) service {
	auth := basic.NewCredentialsBuilder().
		WithAk(cfg.AccessKey).
		WithSk(cfg.SecretKey).
		WithProjectId(cfg.Project).
		Build()

	client := v2.NewNlpClient(core.NewHcHttpClientBuilder().
		WithCredential(auth).
		WithRegion(region.NewRegion(cfg.Region, cfg.AuthEndpoint)).
		Build())

	return service{
		cli: client,
		to:  model.GetTextTranslationReqToEnum(),
	}
}

type service struct {
	cli *v2.NlpClient
	to  model.TextTranslationReqToEnum
}

func (s service) reqTo(to string) model.TextTranslationReqTo {
	switch to {
	case "chinese":
		return s.to.ZH
	case "english":
		return s.to.EN
	default:
		return s.to.EN
	}
}

func (s service) Translate(content string, l dp.Language) (string, error) {
	t := model.TextTranslationReq{
		Text: content,
		From: model.GetTextTranslationReqFromEnum().AUTO,
		To:   s.reqTo(l.Language()),
	}

	req := model.RunTextTranslationRequest{Body: &t}

	v, err := s.cli.RunTextTranslation(&req)
	if err != nil {
		return "", err
	}

	if v.ErrorMsg != nil {
		err = errors.New(*v.ErrorMsg)
		return "", err
	}

	if v.TranslatedText != nil {
		return *v.TranslatedText, nil
	}

	return "", errors.New("no translated text")
}
