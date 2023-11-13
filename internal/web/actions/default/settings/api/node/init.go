package node

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
			Helper(settingutils.NewAdvancedHelper("apiNodes")).
			Prefix("/settings/api/node").

			// 这里不受Helper的约束
			GetPost("/createAddrPopup", new(CreateAddrPopupAction)).
			GetPost("/updateAddrPopup", new(UpdateAddrPopupAction)).

			// 节点相关
			Helper(NewHelper()).
			Get("", new(IndexAction)).
			GetPost("/update", new(UpdateAction)).
			Get("/install", new(InstallAction)).
			Get("/logs", new(LogsAction)).
			GetPost("/upgradePopup", new(UpgradePopupAction)).
			Post("/upgradeCheck", new(UpgradeCheckAction)).

			//
			EndAll()
	})
}
