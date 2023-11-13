// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package reverseProxy

import (
	"github.com/iwind/TeaGo/types"
	teaconst "github.com/oy1978/EdgeAdmin/internal/const"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	GroupId int64
}) {
	if teaconst.IsPlus {
		this.RedirectURL("/servers/groups/group/settings/web?groupId=" + types.String(params.GroupId))
	} else {
		this.RedirectURL("/servers/groups/group/settings/httpReverseProxy?groupId=" + types.String(params.GroupId))
	}
}
