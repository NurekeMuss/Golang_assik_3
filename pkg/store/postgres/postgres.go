package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Port     int
	Host     string
	DbName   string
	Password string
	User     string
}

func NewDBConfig(host string, port int, user, password, dbname string) *DBConfig {
	return &DBConfig{
		Port:     port,
		Host:     host,
		DbName:   dbname,
		Password: password,
		User:     user,
	}
}

func ConnectToDB(cfg *DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
