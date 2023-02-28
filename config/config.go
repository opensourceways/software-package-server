package config

import (
	"errors"

	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/software-package-server/infrastructure/db"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

func LoadConfig(path string, cfg interface{}) (err error) {
	if err = utils.LoadFromYaml(path, cfg); err != nil {
		return
	}

	if f, ok := cfg.(configSetDefault); ok {
		f.SetDefault()
	}

	if f, ok := cfg.(configValidate); ok {
		if err = f.Validate(); err != nil {
			return
		}
	}

	c, ok := cfg.(*Config)
	if !ok {
		return errors.New("assert config fail")
	}

	if err = db.InitPostgresql(&c.PostgresqlConfig); err != nil {
		return
	}

	dp.Init(&c.DPConfig)

	return nil
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
