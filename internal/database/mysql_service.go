package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"idaman.id/storage/internal/config"
)

func NewMySQLClient(configService config.ConfigService) (*sql.DB, error) {
	u := configService.GetString("DB_MYSQL_USERNAME")
	pw := configService.GetString("DB_MYSQL_PASSWORD")
	h := configService.GetString("DB_MYSQL_HOST")
	po := configService.GetString("DB_MYSQL_PORT")
	dn := configService.GetString("DB_MYSQL_NAME")
	cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", u, pw, h, po, dn)

	db, err := sql.Open("mysql", cs)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
