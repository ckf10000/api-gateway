// Package register
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     register_service.go
* Description:  TODO
* Author:       mfkifhss2023
* CreateDate:   2024-10-20 00:24:36
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package register

import (
	"api_gateway/internal/common/utils/log"
	"api_gateway/internal/infrastructure/middleware/consul"
	"sync"
)

var onceRegister sync.Once // 同步锁，只执行一次

func NewRegisterService() {
	onceRegister.Do(func() {
		client := consul.GetConsulClient()
		agentSer := consul.GetAgentService()
		err := client.Agent().ServiceRegister(agentSer)
		if err != nil {
			log.Logger.Error("register service error:", err)
			panic(err)
		}
		log.Logger.Infof("%s service register success", agentSer.ID)
	})
}
