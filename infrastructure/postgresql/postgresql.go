package postgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/opensourceways/software-package-server/config"
)

var db *gorm.DB

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

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetConnMaxLifetime(cfg.DbLife)
	sqlDB.SetMaxOpenConns(cfg.DbMaxConn)
	sqlDB.SetMaxIdleConns(cfg.DbMaxIdle)

	return
}

func DB() *gorm.DB {
	return db
}
