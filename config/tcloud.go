package config

type tcloud struct {
	Sms      sms      `yaml:"sms"`
	Realname realname `yaml:"realname"`
}

type sms struct {
	SecretId  string `yaml:"secretid"`
	SecretKey string `yaml:"secretkey"`
	AppId     string `yaml:"appid"`
	SignName  string `yaml:"signname"`
	TempId    tempId `yaml:"tempid"`
}

type tempId struct {
	Code  string `yaml:"code"`
	Buy   string `yaml:"buy"`
	Login string `yaml:"login"`
}

type realname struct {
	SecretId  string `yaml:"secretid"`
	SecretKey string `yaml:"secretkey"`
}
