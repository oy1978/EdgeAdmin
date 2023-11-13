// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package stat

import (
	"sort"

	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type DailyRequestsAction struct {
	actionutils.ParentAction
}

func (this *DailyRequestsAction) Init() {
	this.Nav("", "stat", "daily")
	this.SecondMenu("index")
}

func (this *DailyRequestsAction) RunGet(params struct {
	ServerId int64
}) {
	this.Data["serverId"] = params.ServerId

	resp, err := this.RPC().ServerDailyStatRPC().FindLatestServerDailyStats(this.AdminContext(), &pb.FindLatestServerDailyStatsRequest{
		ServerId: params.ServerId,
		Days:     30,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	sort.Slice(resp.Stats, func(i, j int) bool {
		stat1 := resp.Stats[i]
		stat2 := resp.Stats[j]
		return stat1.Day < stat2.Day
	})
	statMaps := []maps.Map{}
	for _, stat := range resp.Stats {
		statMaps = append(statMaps, maps.Map{
			"day":                 stat.Day[:4] + "-" + stat.Day[4:6] + "-" + stat.Day[6:8],
			"bytes":               stat.Bytes,
			"cachedBytes":         stat.CachedBytes,
			"countRequests":       stat.CountRequests,
			"countCachedRequests": stat.CountCachedRequests,
		})
	}
	this.Data["dailyStats"] = statMaps

	this.Show()
}
