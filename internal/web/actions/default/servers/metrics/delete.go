// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package metrics

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	ItemId int64
}) {
	defer this.CreateLogInfo(codes.MetricItem_LogDeleteMetricItem)

	_, err := this.RPC().MetricItemRPC().DeleteMetricItem(this.AdminContext(), &pb.DeleteMetricItemRequest{MetricItemId: params.ItemId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
