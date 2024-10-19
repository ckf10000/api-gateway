// Package auth
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     auth.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 11:21:09
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package auth

import (
	"api_gateway/internal/common/utils/log"
	"api_gateway/internal/domain/service/cache"
	"api_gateway/internal/domain/service/jwt"
	"api_gateway/internal/domain/service/user"
	"errors"
	"github.com/gogf/gf/frame/g"
	"time"
)

func Login(username, password string) (map[string]interface{}, error) {
	userDto, err1 := user.GetUserByPhone(username)
	if err1 != nil {
		log.Logger.Error("get user failed: ", err1)
		return nil, err1
	}
	if userDto.PhoneNo != username || userDto.EncryptedPsw != password {
		errStr := "incorrect username or password"
		log.Logger.Error(errStr)
		return nil, errors.New(errStr)
	}
	if !userDto.Active {
		errStr := "user has been deactivated"
		log.Logger.Error(errStr)
		return nil, errors.New(errStr)
	}

	expire, token, err2 := jwt.GenerateToken(username)
	if err2 != nil {
		log.Logger.Error("generate token err: ", err2)
		return nil, err2
	}
	// 将 token 存入 Redis
	err3 := cache.SetCache(username, token, time.Duration(jwt.TokenExp)*time.Hour)
	if err3 != nil {
		log.Logger.Error("token save failed: ", err3)
		return nil, err3
	}
	return g.Map{"token": token, "expire": expire}, nil
}

func Logout(tokenString string) error {
	// 验证并解析 Token
	userID, err := jwt.ValidateToken(tokenString)
	if err != nil {
		log.Logger.Error("invalid token: ", err)
		return err
	}

	// 从 Redis 中删除 Token
	err = cache.DeleteCache(userID)
	if err != nil {
		log.Logger.Error("token delete failed: ", err)
		return err
	}
	return nil
}

func ValidateToken(tokenString string) error {
	// 验证并解析 Token
	userID, err1 := jwt.ValidateToken(tokenString)
	if err1 != nil {
		log.Logger.Error("validate token failed: ", err1)
		return err1
	}
	// 从 Redis 中获取 Token
	_, err2 := cache.GetCache(userID)
	if err2 != nil {
		log.Logger.Error("get token failed: ", err2)
		return err2
	}
	return nil
}
