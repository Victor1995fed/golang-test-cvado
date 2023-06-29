package config

// Config Конфиги приложения
type Config struct {
	App struct {
		Port int `env:"APP_PORT"`
	}
	Database struct {
		Host     string `env:"MYSQL_HOST"`
		Port     string `env:"MYSQL_TCP_PORT"`
		Database string `env:"MYSQL_DATABASE"`
		Username string `env:"DB_USERNAME"`
		Password string `env:"MYSQL_ROOT_PASSWORD"`
	}
}
