package database

import (
	"github.com/iwind/TeaGo"
	"github.com/oy1978/EdgeAdmin/internal/configloaders"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeSetting)).
			Helper(settingutils.NewAdvancedHelper("database")).
			Prefix("/settings/database").
			Get("", new(IndexAction)).
			GetPost("/update", new(UpdateAction)).
			GetPost("/clean", new(CleanAction)).
			GetPost("/cleanSetting", new(CleanSettingAction)).
			GetPost("/truncateTable", new(TruncateTableAction)).
			GetPost("/deleteTable", new(DeleteTableAction)).
			EndAll()
	})
}
