package initialization

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"nft_platform/global"
)

func initRedis() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.Conf.Redis.Host, global.Conf.Redis.Port),
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.Db,
		//DialTimeout:     time.Duration(global.Conf.Redis.Timeout) * time.Second,
		//ReadTimeout:     time.Duration(global.Conf.Redis.Timeout) * time.Second,
		//WriteTimeout:    time.Duration(global.Conf.Redis.Timeout) * time.Second,
		//PoolSize:        global.Conf.Redis.PoolSize,
		//MinIdleConns:    global.Conf.Redis.MinIdle,
		//MaxIdleConns:    global.Conf.Redis.MaxIdle,
		//ConnMaxIdleTime: 60 * time.Second,
		//ConnMaxLifetime: 0,
	})
}
