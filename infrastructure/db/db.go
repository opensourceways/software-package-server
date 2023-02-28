package db

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	sqlDb *sql.DB
	db    *gorm.DB
)

func InitPostgresql(cfg *PostgresqlConfig) (err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.dsn(),
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
