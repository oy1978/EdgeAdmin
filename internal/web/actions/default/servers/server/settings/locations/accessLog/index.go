package accessLog

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
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
	// 获取配置
	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithLocationId(this.AdminContext(), params.LocationId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	this.Data["webId"] = webConfig.Id
	this.Data["accessLogConfig"] = webConfig.AccessLogRef

	// 通用变量
	this.Data["fields"] = serverconfigs.HTTPAccessLogFields
	this.Data["defaultFieldCodes"] = serverconfigs.HTTPAccessLogDefaultFieldsCodes

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId         int64
	AccessLogJSON []byte

	Must *actions.Must
}) {
	defer this.CreateLogInfo(codes.ServerAccessLog_LogUpdateAccessLogSetting, params.WebId)

	// TODO 检查参数

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWebAccessLog(this.AdminContext(), &pb.UpdateHTTPWebAccessLogRequest{
		HttpWebId:     params.WebId,
		AccessLogJSON: params.AccessLogJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
