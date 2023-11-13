// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package users

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type VerifyPopupAction struct {
	actionutils.ParentAction
}

func (this *VerifyPopupAction) RunGet(params struct {
	UserId int64
}) {
	this.Data["userId"] = params.UserId

	this.Show()
}

func (this *VerifyPopupAction) RunPost(params struct {
	UserId       int64
	Result       string
	RejectReason string

	Must *actions.Must
	CSRF *actionutils.CSRF
}) {
	defer this.CreateLogInfo(codes.User_LogVerifyUser, params.UserId, params.Result)

	if params.Result == "pass" {
		params.RejectReason = ""
	}

	_, err := this.RPC().UserRPC().VerifyUser(this.AdminContext(), &pb.VerifyUserRequest{
		UserId:       params.UserId,
		IsRejected:   params.Result == "reject" || params.Result == "delete",
		RejectReason: params.RejectReason,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	if params.Result == "delete" {
		_, err = this.RPC().UserRPC().DeleteUser(this.AdminContext(), &pb.DeleteUserRequest{UserId: params.UserId})
		if err != nil {
			this.ErrorPage(err)
			return
		}
	}

	this.Success()
}
