package regions

import (
	"github.com/iwind/TeaGo"
	"github.com/oy1978/EdgeAdmin/internal/configloaders"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth(configloaders.AdminModuleCodeNode)).
			Data("teaMenu", "clusters").
			Data("teaSubMenu", "region").
			Prefix("/clusters/regions").
			Get("", new(IndexAction)).
			GetPost("/createPopup", new(CreatePopupAction)).
			GetPost("/updatePopup", new(UpdatePopupAction)).
			Post("/delete", new(DeleteAction)).
			Post("/sort", new(SortAction)).
			Get("/nodes", new(NodesAction)).
			GetPost("/updateNodeRegionPopup", new(UpdateNodeRegionPopupAction)).

			//
			GetPost("/selectPopup", new(SelectPopupAction)).

			//
			EndAll()
	})
}
