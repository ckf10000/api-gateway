// Package cache
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     cache_service.go
* Description:  TODO
* Author:       mfkifhss2023
* CreateDate:   2024-10-19 23:31:41
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cache

import (
	Redis "api_gateway/internal/infrastructure/middleware/redis"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

func SetCache(key string, value string, exp time.Duration) error {
	redisClient := Redis.GetRedisClient()
	return redisClient.Client.Set(context.Background(), key, value, exp).Err()
}

func GetCache(key string) (string, error) {
	redisClient := Redis.GetRedisClient()
	value, err := redisClient.Client.Get(context.Background(), key).Result()
	if errors.Is(redis.Nil, err) {
		// Key 不存在或已过期
		return "", errors.New("key does not exist or has expired")
	} else if err != nil {
		// 处理其他错误
		return "", err
	} else {
		// Key 存在且有效
		return value, nil
	}
}

func DeleteCache(key string) error {
	redisClient := Redis.GetRedisClient()
	_, err := redisClient.Client.Del(context.Background(), key).Result()
	return err
}
