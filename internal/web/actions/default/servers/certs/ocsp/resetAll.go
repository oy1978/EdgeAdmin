// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ocsp

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type ResetAllAction struct {
	actionutils.ParentAction
}

func (this *ResetAllAction) RunPost(params struct{}) {
	defer this.CreateLogInfo(codes.SSLCert_LogOCSPResetAllOCSPStatus)

	_, err := this.RPC().SSLCertRPC().ResetAllSSLCertsWithOCSPError(this.AdminContext(), &pb.ResetAllSSLCertsWithOCSPErrorRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
