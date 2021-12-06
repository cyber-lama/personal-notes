package apiserver

type Config struct {
	logLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		logLevel: "debug",
	}
}
