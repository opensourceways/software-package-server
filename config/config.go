package config

import (
	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/software-package-server/infrastructure/db"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

func LoadConfig(path string) (*Config, error) {
	cfg := new(Config)
	if err := utils.LoadFromYaml(path, cfg); err != nil {
		return nil, err
	}

	cfg.SetDefault()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	if err := db.InitPostgresql(&cfg.PostgresqlConfig); err != nil {
		return nil, err
	}

	dp.Init(&cfg.DPConfig)

	return cfg, nil
}

type configValidate interface {
	Validate() error
}

type configSetDefault interface {
	SetDefault()
}

type Config struct {
	PostgresqlConfig db.PostgresqlConfig `json:"db_config" required:"true"`
	DPConfig         dp.Config           `json:"dp_config" required:"true"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.PostgresqlConfig,
		&cfg.DPConfig,
	}
}

func (cfg *Config) SetDefault() {
	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configSetDefault); ok {
			f.SetDefault()
		}
	}
}

func (cfg *Config) Validate() error {
	if _, err := utils.BuildRequestBody(cfg, ""); err != nil {
		return err
	}

	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configValidate); ok {
			if err := f.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
