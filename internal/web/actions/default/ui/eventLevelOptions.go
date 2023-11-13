package ui

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

type EventLevelOptionsAction struct {
	actionutils.ParentAction
}

func (this *EventLevelOptionsAction) RunPost(params struct{}) {
	this.Data["eventLevels"] = firewallconfigs.FindAllFirewallEventLevels()

	this.Success()
}
