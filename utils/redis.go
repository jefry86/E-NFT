package utils

import (
	"context"
	"nft_platform/global"
	"time"
)

var ctx = context.Background()

func RedisGetString(key string) (string, error) {
	val, err := global.Rdb.Get(ctx, key).Result()
	if err != nil && err.Error() != "redis: nil" {
		global.SLogger.Errorf("redis get %s, err:%s", key, err)
		return "", err
	}
	return val, nil
}

func RedisSetString(key, val string, expiration time.Duration) error {
	if err := global.Rdb.Set(ctx, key, val, expiration).Err(); err != nil {
		global.SLogger.Errorf("redis set %s,val:%s err:%s", key, val, err)
		return err
	}
	return nil
}

func RedisInc(key string) error {
	if err := global.Rdb.Incr(ctx, key).Err(); err != nil {
		global.SLogger.Errorf("redis Incr %s err:%s", key, err)
		return err
	}
	return nil
}

func RedisExists(key string) (bool, error) {
	n, err := global.Rdb.Exists(ctx, key).Result()
	if err != nil {
		global.SLogger.Errorf("redis Exists %s err:%s", key, err)
		return false, err
	}
	return n == 1, nil
}

func RedisUnlink(key string) error {
	return global.Rdb.Unlink(ctx, key).Err()
}
