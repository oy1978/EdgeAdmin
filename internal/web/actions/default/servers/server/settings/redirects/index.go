package redirects

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
	this.Nav("", "setting", "index")
	this.SecondMenu("redirects")
}

func (this *IndexAction) RunGet(params struct {
	ServerId int64
}) {
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerId(this.AdminContext(), params.ServerId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id

	this.Data["redirects"] = webConfig.HostRedirects

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	ServerId          int64
	WebId             int64
	HostRedirectsJSON []byte

	Must *actions.Must
}) {
	defer this.CreateLogInfo(codes.ServerRedirect_LogUpdateRedirects, params.WebId)

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWebHostRedirects(this.AdminContext(), &pb.UpdateHTTPWebHostRedirectsRequest{
		HttpWebId:         params.WebId,
		HostRedirectsJSON: params.HostRedirectsJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
