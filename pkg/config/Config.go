package config

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Connection   string
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
}

type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

func NewConfig(server *ServerConfig, database *DatabaseConfig) *Config {
	return &Config{Server: server, Database: database}
}
