package conds

import (
	"github.com/iwind/TeaGo"
	"github.com/oy1978/EdgeAdmin/internal/configloaders"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeServer)).
			Prefix("/servers/server/settings/conds").
			GetPost("/addGroupPopup", new(AddGroupPopupAction)).
			GetPost("/addCondPopup", new(AddCondPopupAction)).
			EndAll()
	})
}
