package initialization

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"nft_platform/global"
)

func initViper() {
	var e string
	flag.StringVar(&e, "e", "dev", "运行环境: dev|qa|prod")
	c := fmt.Sprintf("./resource/conf/application_%s.yml", e)
	v := viper.New()
	v.SetConfigFile(c)
	//v.SetConfigName(fmt.Sprintf("application_%s", e)) //设置读取的文件名
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置出错了,err:%s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.Conf); err != nil {
			panic(fmt.Errorf("Unmarshal conf 出错了,err:%s \n", err))
		}
	})

	if err := v.Unmarshal(&global.Conf); err != nil {
		panic(fmt.Errorf("Unmarshal conf 出错了,err:%s \n", err))
	}
}
