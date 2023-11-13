package db

import (
	"net/http"

	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
)

type Helper struct {
	helpers.LangHelper
}

func (this *Helper) BeforeAction(action *actions.ActionObject) {
	if action.Request.Method != http.MethodGet {
		return
	}

	action.Data["teaMenu"] = "db"

	var selectedTabbar = action.Data["mainTab"]

	var tabbar = actionutils.NewTabbar()
	tabbar.Add(this.Lang(action, codes.DBNode_TabNodes), "", "/db", "", selectedTabbar == "db")
	actionutils.SetTabbar(action, tabbar)
}
