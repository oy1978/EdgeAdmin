package certs

import (
	"encoding/json"
	"strconv"

	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/sslconfigs"
)

type DownloadKeyAction struct {
	actionutils.ParentAction
}

func (this *DownloadKeyAction) Init() {
	this.Nav("", "", "")
}

func (this *DownloadKeyAction) RunGet(params struct {
	CertId int64
}) {
	defer this.CreateLogInfo(codes.SSLCert_LogDownloadSSLCertKey, params.CertId)

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

	this.AddHeader("Content-Disposition", "attachment; filename=\"key-"+strconv.FormatInt(params.CertId, 10)+".pem\";")
	_, _ = this.Write(certConfig.KeyData)
}
