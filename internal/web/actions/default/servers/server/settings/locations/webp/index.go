// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package webp

import (
	"encoding/json"

	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/dao"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {

}

func (this *IndexAction) RunGet(params struct {
	LocationId int64
}) {
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithLocationId(this.AdminContext(), params.LocationId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["webpConfig"] = webConfig.WebP

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId    int64
	WebpJSON []byte

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	var webpConfig = &serverconfigs.WebPImageConfig{}
	err := json.Unmarshal(params.WebpJSON, webpConfig)
	if err != nil {
		this.Fail("参数校验失败：" + err.Error())
	}

	_, err = this.RPC().HTTPWebRPC().UpdateHTTPWebWebP(this.AdminContext(), &pb.UpdateHTTPWebWebPRequest{
		HttpWebId: params.WebId,
		WebpJSON:  params.WebpJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
