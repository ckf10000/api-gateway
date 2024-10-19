// Package cmd
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     manager.go
* Description:  TODO
* Author:       mfkifhss2023
* CreateDate:   2024-10-19 21:17:22
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cmd

import (
	"api_gateway/internal/api/router"
	"api_gateway/internal/application/manager"
	"github.com/gogf/gf/frame/g"
)

func main() {
	// 注册服务到 注册中心
	manager.RegisterService()

	// 启动http服务
	s := g.Server()
	router.RegisterAuthRouter(s)
}
