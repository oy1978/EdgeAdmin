// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package db

import (
	"strings"

	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/default/db/dbnodeutils"
)

type NodeAction struct {
	actionutils.ParentAction
}

func (this *NodeAction) Init() {
	this.Nav("", "", "node")
}

func (this *NodeAction) RunGet(params struct {
	NodeId int64
}) {
	node, err := dbnodeutils.InitNode(this.Parent(), params.NodeId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["node"] = maps.Map{
		"id":          node.Id,
		"isOn":        node.IsOn,
		"name":        node.Name,
		"database":    node.Database,
		"host":        node.Host,
		"port":        node.Port,
		"username":    node.Username,
		"password":    strings.Repeat("*", len(node.Password)),
		"description": node.Description,
	}

	this.Show()
}
