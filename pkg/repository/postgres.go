package repository

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	fmt.Printf("host= %s\n port= %s\n username= %s\n password= %s\n dbname= %s\n sslmode= %s\n", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLMode)
	db, err := sqlx.Open("postgres", fmt.Sprintf("host= %s port= %s username= %s password= %s dbname= %s sslmode= %s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil

}
