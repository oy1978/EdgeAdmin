package charset

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/default/servers/groups/group/servergrouputils"
	"github.com/oy1978/EdgeCommon/pkg/configutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/dao"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "setting", "index")
	this.SecondMenu("charset")
}

func (this *IndexAction) RunGet(params struct {
	GroupId int64
}) {
	_, err := servergrouputils.InitGroup(this.Parent(), params.GroupId, "charset")
	if err != nil {
		this.ErrorPage(err)
		return
	}

	webConfig, err := dao.SharedHTTPWebDAO.FindWebConfigWithServerGroupId(this.AdminContext(), params.GroupId)
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Data["webId"] = webConfig.Id
	this.Data["charsetConfig"] = webConfig.Charset

	this.Data["usualCharsets"] = configutils.UsualCharsets
	this.Data["allCharsets"] = configutils.AllCharsets

	this.Show()
}

func (this *IndexAction) RunPost(params struct {
	WebId       int64
	CharsetJSON []byte

	Must *actions.Must
}) {
	defer this.CreateLogInfo(codes.ServerCharset_LogUpdateCharsetSetting, params.WebId)

	_, err := this.RPC().HTTPWebRPC().UpdateHTTPWebCharset(this.AdminContext(), &pb.UpdateHTTPWebCharsetRequest{
		HttpWebId:   params.WebId,
		CharsetJSON: params.CharsetJSON,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
