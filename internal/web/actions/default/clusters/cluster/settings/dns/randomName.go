// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dns

import (
	"github.com/iwind/TeaGo/rands"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
)

type RandomNameAction struct {
	actionutils.ParentAction
}

func (this *RandomNameAction) RunPost(params struct{}) {
	this.Data["name"] = "cluster" + rands.HexString(8)

	this.Success()
}
