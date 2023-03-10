package translationimpl

type Config struct {
	TranslationURL string      `json:"translation_url" required:"true"`
	Cloud          CloudConfig `json:"cloud"           required:"true"`
}

type CloudConfig struct {
	Domain       string `json:"domain"                 required:"true"`
	User         string `json:"user"                   required:"true"`
	Password     string `json:"password"               required:"true"`
	Project      string `json:"project"                required:"true"`
	AuthEndpoint string `json:"auth_endpoint"          required:"true"`
}
