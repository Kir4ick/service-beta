package config

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host string
	Port string
}

type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

func NewConfig(server *ServerConfig, database *DatabaseConfig) *Config {
	return &Config{Server: server, Database: database}
}
