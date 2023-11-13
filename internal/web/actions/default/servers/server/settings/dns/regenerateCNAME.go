// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dns

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type RegenerateCNAMEAction struct {
	actionutils.ParentAction
}

func (this *RegenerateCNAMEAction) RunPost(params struct {
	ServerId int64
}) {
	defer this.CreateLogInfo(codes.ServerDNS_LogRegenerateDNSName, params.ServerId)

	_, err := this.RPC().ServerRPC().RegenerateServerDNSName(this.AdminContext(), &pb.RegenerateServerDNSNameRequest{ServerId: params.ServerId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
