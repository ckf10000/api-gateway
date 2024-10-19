// Package auth
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     auth_service.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 11:21:09
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package auth

import (
	"api_gateway/internal/common/utils/log"
	"api_gateway/internal/domain/service/jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Login(r *ghttp.Request) {
	username := r.GetString("username")
	password := r.GetString("password")

	if username == "user" || password == "pass" {
		token, err1 := jwt.GenerateToken(username)
		if err1 != nil {
			log.Logger.Error("Generate token err:", err1)
			r.Response.WriteStatus(http.StatusInternalServerError)
			return
		}
		err2 := r.Response.WriteJson(g.Map{"code": "200", "data": g.Map{"token": token}, "message": "Generate token successfully", "success": true})
		if err2 != nil {
			r.Response.WriteStatus(http.StatusInternalServerError)
			return
		}
	} else {
		r.Response.WriteHeader(http.StatusUnauthorized)
	}
}

func Logout(r *ghttp.Request) {
	err := r.Response.WriteJson(g.Map{"code": "200", "message": "Logged out successfully", "data": nil, "success": true})
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}
}
