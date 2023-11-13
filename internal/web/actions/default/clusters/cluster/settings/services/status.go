package services

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/default/nodes/nodeutils"
	"github.com/oy1978/EdgeCommon/pkg/messageconfigs"
)

type StatusAction struct {
	actionutils.ParentAction
}

func (this *StatusAction) Init() {
	this.Nav("", "setting", "status")
	this.SecondMenu("service")
}

func (this *StatusAction) RunGet(params struct {
}) {
	this.Show()
}

func (this *StatusAction) RunPost(params struct {
	ClusterId int64

	Must *actions.Must
}) {
	results, err := nodeutils.SendMessageToCluster(this.AdminContext(), params.ClusterId, messageconfigs.MessageCodeCheckSystemdService, &messageconfigs.CheckSystemdServiceMessage{}, 10)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["results"] = results

	this.Success()
}
