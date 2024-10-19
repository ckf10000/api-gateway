// Package consul
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     consul.go
* Description:  TODO
* Author:       mfkifhss2023
* CreateDate:   2024-10-19 23:37:23
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package consul

import (
	"api_gateway/internal/common/utils/log"
	"github.com/hashicorp/consul/api"
)

func GetConsulClient() *api.Client {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Logger.Error("create consul client error:", err)
		panic(err)
	}
	return client
}

func GetAgentService() *api.AgentServiceRegistration {
	return &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: address,
		Port:    port,
	}
}
