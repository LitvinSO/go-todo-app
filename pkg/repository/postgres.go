package repository

import (
	"fmt"
	"net/url"

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
	TimeZone string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	//fmt.Printf("host= %s\n port= %s\n username= %s\n password= %s\n dbname= %s\n sslmode= %s\n", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLMode)
	//db, err := sqlx.Open("postgres", fmt.Sprintf("host= %s port= %s username= %s password= %s dbname= %s sslmode= %s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLMode))
	q := url.Values{}
	q.Set("sslmode", cfg.SSLMode)
	q.Set("timezone", cfg.TimeZone)

	u := url.URL{
		Scheme:   cfg.DBname,
		User:     url.UserPassword(cfg.Username, cfg.Password),
		Host:     cfg.Host + ":" + cfg.Port,
		Path:     "postgres",
		RawQuery: q.Encode(),
	}

	fmt.Println(u.String())

	db, err := sqlx.Open("postgres", u.String())

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil

}
