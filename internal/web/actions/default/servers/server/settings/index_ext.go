// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package settings

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

func (this *IndexAction) initUserPlan(server *pb.Server) {
	var userPlanMap = maps.Map{"id": server.UserPlanId, "dayTo": "", "plan": maps.Map{}}
	this.Data["userPlan"] = userPlanMap
}
