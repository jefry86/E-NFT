package config

type server struct {
	Name              string `yaml:"name"`
	Port              string `yaml:"port"`
	Mode              string `yaml:"mode"`
	AppKey            string `yaml:"appkey"`
	AppSecret         string `yaml:"appsecret"`
	SmsMobileLimit    int    `yaml:"smsmobilelimit"`
	SmsIPLimit        int    `yaml:"smsiplimit"`
	DefaultAvatar     string `yaml:"defaultavatar"`
	PayExpirationTime int    `yaml:"payexpirationtime"`
}
