package ipadmin

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type DeleteIPAction struct {
	actionutils.ParentAction
}

func (this *DeleteIPAction) RunPost(params struct {
	FirewallPolicyId int64
	ItemId           int64
}) {
	// 日志
	defer this.CreateLogInfo(codes.WAF_LogDeleteIPFromWAFPolicy, params.FirewallPolicyId, params.ItemId)

	// TODO 判断权限

	_, err := this.RPC().IPItemRPC().DeleteIPItem(this.AdminContext(), &pb.DeleteIPItemRequest{IpItemId: params.ItemId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
