package http

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/dao"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
}

func (this *IndexAction) RunGet(params struct {
	LocationId int64
}) {
	// 跳转相关设置
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithLocationId(this.AdminContext(), params.LocationId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["webId"] = webConfig.Id
	this.Data["redirectToHTTPSConfig"] = webConfig.RedirectToHttps

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId               int64
	RedirectToHTTPSJSON []byte

	Must *actions.Must
}) {
	defer this.CreateLogInfo(codes.ServerRedirect_LogUpdateRedirects, params.WebId)

	// 设置跳转到HTTPS
	// TODO 校验设置
	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWebRedirectToHTTPS(this.AdminContext(), &pb.UpdateHTTPWebRedirectToHTTPSRequest{
		HttpWebId:           params.WebId,
		RedirectToHTTPSJSON: params.RedirectToHTTPSJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
