// Package redis
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     redis.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 14:24:52
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package redis

import (
	"api_gateway/internal/common/utils/log"
	"github.com/go-redis/redis/v8"
	"sync"
)

type RedClient struct {
	Client *redis.Client
}

var (
	redisClient *RedClient // redis客户端
	onceRedis   sync.Once  // 同步锁，只执行一次
)

func GetRedisClient() *RedClient {
	onceRedis.Do(func() {
		redisClient = &RedClient{
			Client: redis.NewClient(&redis.Options{
				Addr:     addr,
				DB:       db,
				Password: password,
				PoolSize: poolSize,
			}),
		}
		log.Logger.Info("redis connect success")
	})
	return redisClient
}
