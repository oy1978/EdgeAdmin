// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package recovers

import "github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct{}) {
	this.Show()
}
