package config

type AppConfig struct {
	Mode   string
	Server struct {
		Host string
		Port uint16
	}
	Database struct {
		Dialect  string
		Host     string
		Port     uint16
		Name     string
		User     string
		Password string
	}
	JWT struct {
		Secret string
	}
}

var Const AppConfig
