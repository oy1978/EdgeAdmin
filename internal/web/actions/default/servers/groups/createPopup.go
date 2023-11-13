package groups

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type CreatePopupAction struct {
	actionutils.ParentAction
}

func (this *CreatePopupAction) Init() {
	this.Nav("", "", "")
}

func (this *CreatePopupAction) RunGet(params struct{}) {
	this.Show()
}

func (this *CreatePopupAction) RunPost(params struct {
	Name string

	Must *actions.Must
}) {
	params.Must.
		Field("name", params.Name).
		Require("请输入分组名称")
	createResp, err := this.RPC().ServerGroupRPC().CreateServerGroup(this.AdminContext(), &pb.CreateServerGroupRequest{
		Name: params.Name,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["group"] = maps.Map{
		"id":   createResp.ServerGroupId,
		"name": params.Name,
	}

	// 创建日志
	defer this.CreateLogInfo(codes.ServerGroup_LogCreateServerGroup, createResp.ServerGroupId)

	this.Success()
}
