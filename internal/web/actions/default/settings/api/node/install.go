package node

import (
	"os"

	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type InstallAction struct {
	actionutils.ParentAction
}

func (this *InstallAction) Init() {
	this.Nav("", "", "install")
}

func (this *InstallAction) RunGet(params struct {
	NodeId int64
}) {
	// API节点信息
	nodeResp, err := this.RPC().APINodeRPC().FindEnabledAPINode(this.AdminContext(), &pb.FindEnabledAPINodeRequest{ApiNodeId: params.NodeId})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	node := nodeResp.ApiNode
	if node == nil {
		this.NotFound("apiNode", params.NodeId)
		return
	}

	this.Data["node"] = maps.Map{
		"id":       node.Id,
		"name":     node.Name,
		"uniqueId": node.UniqueId,
		"secret":   node.Secret,
	}

	// 数据库配置
	dbConfigMap := maps.Map{
		"config":     "",
		"error":      "",
		"isNotFound": false,
	}
	dbConfigFile := Tea.ConfigFile("api_db.yaml")
	data, err := os.ReadFile(dbConfigFile)
	dbConfigMap["config"] = string(data)
	if err != nil {
		dbConfigMap["error"] = err.Error()
		dbConfigMap["isNotFound"] = os.IsNotExist(err)
	}
	this.Data["dbConfig"] = dbConfigMap

	this.Show()
}
