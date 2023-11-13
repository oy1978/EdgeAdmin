package certs

import (
	"encoding/json"

	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

type ViewKeyAction struct {
	actionutils.ParentAction
}

func (this *ViewKeyAction) Init() {
	this.Nav("", "", "")
}

func (this *ViewKeyAction) RunGet(params struct {
	CertId int64
}) {
	certResp, err := this.RPC().SSLCertRPC().FindEnabledSSLCertConfig(this.AdminContext(), &pb.FindEnabledSSLCertConfigRequest{SslCertId: params.CertId})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	certConfig := &sslconfigs.SSLCertConfig{}
	err = json.Unmarshal(certResp.SslCertJSON, certConfig)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	_, _ = this.Write(certConfig.KeyData)
}
