package waf

import "github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"

type TestAction struct {
	actionutils.ParentAction
}

func (this *TestAction) Init() {
	this.Nav("", "", "test")
}

func (this *TestAction) RunGet(params struct{}) {
	this.Show()
}
