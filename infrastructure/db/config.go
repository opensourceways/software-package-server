package db

import (
	"fmt"
	"time"
)

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

func (p *PostgresqlConfig) DSN() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		p.DbHost, p.DbUser, p.DbPwd, p.DbName, p.DbPort)
}
