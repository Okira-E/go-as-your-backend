package system_users

import "github.com/jmoiron/sqlx"

var config Config

type Config struct {
	datasource *sqlx.DB
}

func InitConfig(db *sqlx.DB) {
	config.datasource = db
}