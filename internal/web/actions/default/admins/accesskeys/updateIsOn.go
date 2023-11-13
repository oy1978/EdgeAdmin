package accesskeys

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type UpdateIsOnAction struct {
	actionutils.ParentAction
}

func (this *UpdateIsOnAction) RunPost(params struct {
	AccessKeyId int64
	IsOn        bool
}) {
	defer this.CreateLogInfo(codes.UserAccessKey_LogUpdateUserAccessKeyIsOn, params.AccessKeyId)

	_, err := this.RPC().UserAccessKeyRPC().UpdateUserAccessKeyIsOn(this.AdminContext(), &pb.UpdateUserAccessKeyIsOnRequest{
		UserAccessKeyId: params.AccessKeyId,
		IsOn:            params.IsOn,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
