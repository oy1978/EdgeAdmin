package clusters

import (
	"github.com/iwind/TeaGo"
	"github.com/oy1978/EdgeAdmin/internal/configloaders"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/default/clusters/clusterutils"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNode)).
			Helper(clusterutils.NewClustersHelper()).
			Data("teaMenu", "clusters").
			Data("teaSubMenu", "cluster").
			Prefix("/clusters").
			Get("", new(IndexAction)).
			GetPost("/create", new(CreateAction)).
			GetPost("/createNode", new(CreateNodeAction)).
			Post("/pin", new(PinAction)).
			Get("/nodes", new(NodesAction)).

			// 只要登录即可访问的Action
			EndHelpers().
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeCommon)).
			Post("/options", new(OptionsAction)).
			Post("/nodeOptions", new(NodeOptionsAction)).
			GetPost("/selectPopup", new(SelectPopupAction)).
			EndAll()
	})
}
