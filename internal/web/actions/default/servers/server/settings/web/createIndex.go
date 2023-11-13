package web

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
)

// 创建首页文件
type CreateIndexAction struct {
	actionutils.ParentAction
}

func (this *CreateIndexAction) Init() {
	this.Nav("", "", "")
}

func (this *CreateIndexAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreateIndexAction) RunPost(params struct {
	Index string

	Must *actions.Must
}) {
	params.Must.
		Field("index", params.Index).
		Require("首页文件不能为空")

	this.Data["index"] = params.Index
	this.Success()
}
