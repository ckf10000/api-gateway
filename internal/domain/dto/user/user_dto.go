// Package user
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     user_dto.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 16:07:24
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package user

type QueryUserDTO struct {
	ID           int32  `json:"id"`
	CompanyID    int32  `json:"company_id"`
	AppID        int32  `json:"app_id"`
	PhoneNo      string `json:"phone_no"`
	EncryptedPsw string `json:"encrypted_psw"`
	Active       bool   `json:"active"`
	Name         string `json:"name"`
}
