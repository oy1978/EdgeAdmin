// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ddosProtection

import (
	"encoding/json"

	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/default/nodes/nodeutils"
	"github.com/oy1978/EdgeCommon/pkg/messageconfigs"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/ddosconfigs"
)

type StatusAction struct {
	actionutils.ParentAction
}

func (this *StatusAction) Init() {
	this.Nav("", "setting", "")
	this.SecondMenu("ddosProtection")
}

func (this *StatusAction) RunGet(params struct {
	ClusterId int64
}) {
	this.Data["clusterId"] = params.ClusterId

	this.Show()
}

func (this *StatusAction) RunPost(params struct {
	ClusterId int64
}) {
	results, err := nodeutils.SendMessageToCluster(this.AdminContext(), params.ClusterId, messageconfigs.MessageCodeCheckLocalFirewall, &messageconfigs.CheckLocalFirewallMessage{
		Name: "nftables",
	}, 10)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	var resultMaps = []maps.Map{}
	for _, result := range results {
		var resultMap = maps.Map{
			"isOk":     result.IsOK,
			"message":  result.Message,
			"nodeId":   result.NodeId,
			"nodeName": result.NodeName,
		}

		nodeResp, err := this.RPC().NodeRPC().FindNodeDDoSProtection(this.AdminContext(), &pb.FindNodeDDoSProtectionRequest{NodeId: result.NodeId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		if len(nodeResp.DdosProtectionJSON) > 0 {
			var ddosProtection = ddosconfigs.DefaultProtectionConfig()
			err = json.Unmarshal(nodeResp.DdosProtectionJSON, ddosProtection)
			if err != nil {
				this.ErrorPage(err)
				return
			}
			resultMap["isPrior"] = !ddosProtection.IsPriorEmpty()
		}
		resultMaps = append(resultMaps, resultMap)
	}

	this.Data["results"] = resultMaps
	this.Success()
}
