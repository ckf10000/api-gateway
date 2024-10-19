// Package log
/*********************************************************************************************************************
* ProjectName:  api_gateway
* FileName:     log.go
* Description:  TODO
* Author:       mfkif
* CreateDate:   2024-10-19 14:47:35
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package log

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var Logger = glog.New()

func init() {
	err := Logger.SetConfigWithMap(g.Map{
		"path":     ".",
		"file":     "api_gateway.log",
		"level":    "all",
		"stdout":   true,
		"StStatus": 0,
	})
	if err != nil {
		return
	}
}
