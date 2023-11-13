package waf

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type UpdateSetOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateSetOnAction) RunPost(params struct {
	SetId int64
	IsOn  bool
}) {
	// 日志
	defer this.CreateLogInfo(codes.WAFRuleSet_LogUpdateRuleSetIsOn, params.SetId)

	_, err := this.RPC().HTTPFirewallRuleSetRPC().UpdateHTTPFirewallRuleSetIsOn(this.AdminContext(), &pb.UpdateHTTPFirewallRuleSetIsOnRequest{
		FirewallRuleSetId: params.SetId,
		IsOn:              params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
