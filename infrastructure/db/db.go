package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/opensourceways/software-package-server/config"
)

var (
	sqlDb *sql.DB
	db    *gorm.DB
)

func InitPostgresql(cfg *config.PostgresqlConfig) (err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DbHost, cfg.DbUser, cfg.DbPwd, cfg.DbName, cfg.DbPort)
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDb, err = db.DB()
	if err != nil {
		return
	}

	sqlDb.SetConnMaxLifetime(cfg.DbLife)
	sqlDb.SetMaxOpenConns(cfg.DbMaxConn)
	sqlDb.SetMaxIdleConns(cfg.DbMaxIdle)

	return
}

func DB() *gorm.DB {
	return db
}
