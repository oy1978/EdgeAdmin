// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package logs

import (
	"github.com/iwind/TeaGo"
	"github.com/oy1978/EdgeAdmin/internal/configloaders"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Data("teaMenu", "servers").
			Data("teaSubMenu", "log").
			Prefix("/servers/logs").
			Get("", new(IndexAction)).
			GetPost("/settings", new(SettingsAction)).
			Post("/partitionData", new(PartitionDataAction)).
			Post("/hasLogs", new(HasLogsAction)).
			EndAll()
	})
}
