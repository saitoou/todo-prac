package config

import (
	"fmt"
	"net/url"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Database database
	Auth     auth0
}

type database struct {
	Scheme   string `env:"DB_SCHEME" envDefault:"postgres"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Database string `env:"DB_DATABASE" envDefault:"todo_db"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

type auth0 struct {
	JWT_SECRET string `env:"AUTO0_JWT_SECRET" envDefault:""`
}

var AppConf *Config

func PostgresURL() string {
	userInfo := url.UserPassword(AppConf.Database.User, AppConf.Database.Password)

	databaseURL := &url.URL{
		Scheme: AppConf.Database.Scheme,
		User:   userInfo,
		Host:   fmt.Sprintf("%s:%s", AppConf.Database.Host, AppConf.Database.Port),
		Path:   AppConf.Database.Database,
	}
	return databaseURL.String()
}

func init() {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
	}
	AppConf = &cfg
}
