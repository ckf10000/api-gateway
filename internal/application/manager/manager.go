// Package manager
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     manager.go
* Description:  TODO
* Author:       mfkifhss2023
* CreateDate:   2024-10-20 00:48:47
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"api_gateway/internal/domain/service/register"
)

func RegisterService() {
	register.NewRegisterService()
}
