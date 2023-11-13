package accesskeys

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type DeleteAction struct {
	actionutils.ParentAction
}

func (this *DeleteAction) RunPost(params struct {
	AccessKeyId int64
}) {
	defer this.CreateLogInfo(codes.UserAccessKey_LogDeleteUserAccessKey, params.AccessKeyId)

	_, err := this.RPC().UserAccessKeyRPC().DeleteUserAccessKey(this.AdminContext(), &pb.DeleteUserAccessKeyRequest{UserAccessKeyId: params.AccessKeyId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
