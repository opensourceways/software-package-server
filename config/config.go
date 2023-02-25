package config

import (
	"time"

	"github.com/opensourceways/community-robot-lib/utils"

	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
)

func LoadConfig(path string, cfg interface{}) error {
	if err := utils.LoadFromYaml(path, cfg); err != nil {
		return err
	}

	if f, ok := cfg.(configSetDefault); ok {
		f.SetDefault()
	}

	if f, ok := cfg.(configValidate); ok {
		if err := f.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type configValidate interface {
	Validate() error
}

type configSetDefault interface {
	SetDefault()
}

type Config struct {
	PostgresqlConfig PostgresqlConfig `json:"postgresql_config" required:"true"`
	DPConfig         dp.Config        `json:"dp_config" required:"true"`
}

type PostgresqlConfig struct {
	DbHost    string        `json:"db_host" required:"true"`
	DbUser    string        `json:"db_user" required:"true"`
	DbPwd     string        `json:"db_pwd"  required:"true"`
	DbName    string        `json:"db_name" required:"true"`
	DbPort    int           `json:"db_port" required:"true"`
	DbMaxConn int           `json:"db_max_conn" required:"true"`
	DbMaxIdle int           `json:"db_max_idle" required:"true"`
	DbLife    time.Duration `json:"db_life" required:"true"`
}

func (p *PostgresqlConfig) SetDefault() {
	if p.DbMaxConn <= 0 {
		p.DbMaxConn = 500
	}

	if p.DbMaxIdle <= 0 {
		p.DbMaxIdle = 250
	}

	if p.DbLife <= 0 {
		p.DbLife = time.Minute * 2
	}
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
