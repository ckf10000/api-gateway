// Package postgres
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     postgres.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 16:13:09
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package postgres

import (
	"api_gateway/internal/common/utils/log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
)

type DBClient struct {
	DB *sql.DB
}

var (
	dbClient *DBClient
	onceDB   sync.Once
)

func GetDBClient() *DBClient {
	onceDB.Do(func() {
		var err error
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s, connect_timeout=%d", host, port, user, password, dbName, sslMode, timeOut)
		dbClient = &DBClient{}
		dbClient.DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Logger.Error(err)
		}
		if err = dbClient.DB.Ping(); err != nil {
			log.Logger.Error(err)
		}
		log.Logger.Info("Successfully connected to postgres database")
	})

	return dbClient
}
