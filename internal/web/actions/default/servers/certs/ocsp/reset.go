// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package ocsp

import (
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type ResetAction struct {
	actionutils.ParentAction
}

func (this *ResetAction) RunPost(params struct {
	CertIds []int64
}) {
	defer this.CreateLogInfo(codes.SSLCert_LogOCSPResetOCSPStatus)

	_, err := this.RPC().SSLCertRPC().ResetSSLCertsWithOCSPError(this.AdminContext(), &pb.ResetSSLCertsWithOCSPErrorRequest{SslCertIds: params.CertIds})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
