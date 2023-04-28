package config

type Config struct {
	Mysql  mysql
	Redis  redis
	Server server
	Logger logger
	TCloud tcloud
}
