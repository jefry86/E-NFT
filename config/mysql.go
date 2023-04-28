package config

type mysql struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	MaxIdle     int    `yaml:"max-idle"`
	MaxOpenIdle int    `yaml:"max-open-idle"`
	MaxLifeTime int    `yaml:"max-life-time"`
}
