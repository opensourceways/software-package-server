package translationimpl

type Config struct {
	AccessKey    string   `json:"access_key"             required:"true"`
	SecretKey    string   `json:"secret_key"             required:"true"`
	Project      string   `json:"project"                required:"true"`
	Region       string   `json:"region"                 required:"true"`
	AuthEndpoint []string `json:"auth_endpoint"          required:"true"`
}
