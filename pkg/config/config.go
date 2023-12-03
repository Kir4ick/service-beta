package config

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
	GammaUrl string
}

func NewConfig(server *ServerConfig, database *DatabaseConfig) *Config {
	return &Config{Server: server, Database: database}
}

func (config *Config) SetGammaUrl(gammaUrl string) *Config {
	config.GammaUrl = gammaUrl
	return config
}
