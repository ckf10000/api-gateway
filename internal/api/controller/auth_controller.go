// Package controller
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     auth_controller.go
* Description:  TODO
* Author:       mfkifhss2023
* CreateDate:   2024-10-20 01:20:52
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package controller

import (
	"api_gateway/internal/application/auth"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func Login(r *ghttp.Request) {
	username := r.GetString("username")
	password := r.GetString("password")

	m, err := auth.Login(username, password)
	var respBody map[string]interface{}

	if err != nil {
		respBody = g.Map{"code": "400", "message": err, "data": nil, "success": false}
	} else {
		respBody = g.Map{"code": "200", "message": "user logged in success", "data": m, "success": true}
	}
	err = r.Response.WriteJson(respBody)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

}

func Logout(r *ghttp.Request) {
	tokenString := r.GetString("token") // 从请求中获取 token

	err := auth.Logout(tokenString)
	var respBody map[string]interface{}
	if err != nil {
		respBody = g.Map{"code": "400", "message": err, "data": nil, "success": false}
	} else {
		respBody = g.Map{"code": "200", "message": "Logged out success", "data": nil, "success": true}
	}
	err = r.Response.WriteJson(respBody)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}
}

func ValidateToken(r *ghttp.Request) {
	tokenString := r.GetString("token")

	err := auth.ValidateToken(tokenString)
	var respBody map[string]interface{}
	if err != nil {
		respBody = g.Map{"code": "400", "message": err, "data": nil, "success": false}
	} else {
		respBody = g.Map{"code": "200", "message": "validate token passed", "data": nil, "success": true}
	}
	err = r.Response.WriteJson(respBody)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}

}
