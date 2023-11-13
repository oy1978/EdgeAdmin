package cache

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

// 计算可用缓存策略数量
type CountAction struct {
	actionutils.ParentAction
}

func (this *CountAction) RunPost(params struct{}) {
	countResp, err := this.RPC().HTTPCachePolicyRPC().CountAllEnabledHTTPCachePolicies(this.AdminContext(), &pb.CountAllEnabledHTTPCachePoliciesRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["count"] = countResp.Count

	this.Success()
}
