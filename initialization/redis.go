package initialization

import (
	"github.com/redis/go-redis/v9"
	"nft_platform/global"
	"time"
)

func initRedis() {
	redis.NewClient(&redis.Options{
		Addr:            global.Conf.Redis.Host,
		Password:        global.Conf.Redis.Password,
		DB:              global.Conf.Redis.Db,
		DialTimeout:     time.Duration(global.Conf.Redis.Timeout) * time.Second,
		ReadTimeout:     time.Duration(global.Conf.Redis.Timeout) * time.Second,
		WriteTimeout:    time.Duration(global.Conf.Redis.Timeout) * time.Second,
		PoolSize:        global.Conf.Redis.PoolSize,
		MinIdleConns:    global.Conf.Redis.MinIdle,
		MaxIdleConns:    global.Conf.Redis.MaxIdle,
		ConnMaxIdleTime: 60 * time.Second,
		ConnMaxLifetime: 0,
	})
}
