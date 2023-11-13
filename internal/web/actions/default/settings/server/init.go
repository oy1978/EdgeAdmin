package server

import (
	"github.com/iwind/TeaGo"
	"github.com/oy1978/EdgeAdmin/internal/configloaders"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Helper(settingutils.NewHelper("server")).
			Prefix("/settings/server").
			Get("", new(IndexAction)).
			GetPost("/updateHTTPPopup", new(UpdateHTTPPopupAction)).
			GetPost("/updateHTTPSPopup", new(UpdateHTTPSPopupAction)).
			EndAll()
	})
}
