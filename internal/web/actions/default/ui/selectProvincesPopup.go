package ui

import (
	"github.com/iwind/TeaGo/actions"
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
	ProvinceIds string
}) {
	var selectedProvinceIds = utils.SplitNumbers(params.ProvinceIds)

	provincesResp, err := this.RPC().RegionProvinceRPC().FindAllRegionProvincesWithRegionCountryId(this.AdminContext(), &pb.FindAllRegionProvincesWithRegionCountryIdRequest{RegionCountryId: regionconfigs.RegionChinaId})
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

func (this *SelectProvincesPopupAction) RunPost(params struct {
	ProvinceIds []int64

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	var provinceMaps = []maps.Map{}
	for _, provinceId := range params.ProvinceIds {
		provinceResp, err := this.RPC().RegionProvinceRPC().FindRegionProvince(this.AdminContext(), &pb.FindRegionProvinceRequest{RegionProvinceId: provinceId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
		province := provinceResp.RegionProvince
		if province == nil {
			continue
		}
		provinceMaps = append(provinceMaps, maps.Map{
			"id":   province.Id,
			"name": province.DisplayName,
		})
	}
	this.Data["provinces"] = provinceMaps

	this.Success()
}
