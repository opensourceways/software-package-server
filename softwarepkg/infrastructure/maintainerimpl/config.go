package maintainerimpl

type Config struct {
	PermissionURL string `json:"permission_url" required:"true"`
	Community     string `json:"community"      required:"true"`
}

func (cfg *Config) SetDefault() {
	if cfg.Community == "" {
		cfg.Community = "openeuler"
	}
}
