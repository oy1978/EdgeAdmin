package dashboard

import (
	"github.com/iwind/TeaGo"
	"github.com/oy1978/EdgeAdmin/internal/configloaders"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.Prefix("/dashboard").
			Data("teaMenu", "dashboard").
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeCommon)).
			GetPost("", new(IndexAction)).
			Post("/restartLocalAPINode", new(RestartLocalAPINodeAction)).
			EndAll()
	})
}
