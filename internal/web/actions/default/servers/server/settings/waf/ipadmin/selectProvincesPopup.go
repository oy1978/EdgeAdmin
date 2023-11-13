// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package ipadmin

import (
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/utils"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/regionconfigs"
)

type SelectProvincesPopupAction struct {
	actionutils.ParentAction
}

func (this *SelectProvincesPopupAction) Init() {
	this.Nav("", "", "")
}

func (this *SelectProvincesPopupAction) RunGet(params struct {
	Type                string
	SelectedProvinceIds string
}) {
	this.Data["type"] = params.Type

	var selectedProvinceIds = utils.SplitNumbers(params.SelectedProvinceIds)

	provincesResp, err := this.RPC().RegionProvinceRPC().FindAllRegionProvincesWithRegionCountryId(this.AdminContext(), &pb.FindAllRegionProvincesWithRegionCountryIdRequest{
		RegionCountryId: regionconfigs.RegionChinaId,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var provinceMaps = []maps.Map{}
	for _, province := range provincesResp.RegionProvinces {
		provinceMaps = append(provinceMaps, maps.Map{
			"id":        province.Id,
			"name":      province.DisplayName,
			"isChecked": lists.ContainsInt64(selectedProvinceIds, province.Id),
		})
	}
	this.Data["provinces"] = provinceMaps

	this.Show()
}
