// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplists

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeAdmin/internal/web/helpers"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type ReadAllAction struct {
	actionutils.ParentAction
}

func (this *ReadAllAction) RunPost(params struct{}) {
	defer this.CreateLogInfo(codes.IPItem_LogReadAllIPItems)

	_, err := this.RPC().IPItemRPC().UpdateIPItemsRead(this.AdminContext(), &pb.UpdateIPItemsReadRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	// 通知左侧菜单Badge更新
	helpers.NotifyIPItemsCountChanges()

	this.Success()
}
