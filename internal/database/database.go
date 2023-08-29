package database

import (
	"context"
	"github.com/avag-sargsyan/gambling_platform/internal/conf"
	"github.com/avag-sargsyan/gambling_platform/internal/util"
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	rdb  *redis.Client
	once sync.Once
	ctx  = context.Background()
)

func GetConnection(configApp *conf.App) *redis.Client {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr: configApp.RedisAddress,
		})

		_, err := rdb.Ping(ctx).Result()
		util.FatalIfError(err)
	})
	return rdb
}
