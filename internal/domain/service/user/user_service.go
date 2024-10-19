// Package user
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     user_service.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 17:50:52
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package user

import (
	"api_gateway/internal/domain/dto/user"
	"api_gateway/internal/infrastructure/database/postgres"
)

func GetUserByPhone(phone string) (*user.QueryUserDTO, error) {
	client := postgres.GetDBClient()
	var u user.QueryUserDTO
	query := `SELECT id, company_id, app_id, phone_no, encrypted_psw, active, name FROM T_user WHERE phone_no = $1`
	err := client.DB.QueryRow(query, phone).Scan(&u.ID, &u.CompanyID, &u.AppID, &u.PhoneNo, &u.EncryptedPsw, &u.Active, &u.Name)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
