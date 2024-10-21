package config

type ConifgService interface {
	GetConfig(...string) *Config
}

type (
	Config struct {
		Server ServerConfig `yaml: server`
		Db     DbConfig     `yaml: db`
	}

	ServerConfig struct {
		Host string `yaml: host`
		Port string `yaml: port`
	}

	DbConfig struct {
		Host     string `yaml: host`
		Port     string `yaml: port`
		Reponame string `yaml: reponame `
		Username string `yaml: username`
		Password string `yaml: password`
	}
)
