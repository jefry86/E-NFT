package config

type redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Timeout  int    `yaml:"timeout"`
	MaxIdle  int    `yaml:"max-idle"`
	MinIdle  int    `yaml:"min-idle"`
	PoolSize int    `yaml:"pool-size"`
}
