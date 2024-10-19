// Package router
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     auth_router.go
* Description:  TODO
* Author:       mfkifhss2023
* CreateDate:   2024-10-19 23:47:01
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package router

import (
	"api_gateway/internal/api/controller"
	"github.com/gogf/gf/net/ghttp"
)

func RegisterAuthRouter(s *ghttp.Server) {
	s.BindHandler("/login", controller.Login)
	s.BindHandler("/logout", controller.Logout)
	s.BindHandler("/validateToken", controller.ValidateToken)
}
