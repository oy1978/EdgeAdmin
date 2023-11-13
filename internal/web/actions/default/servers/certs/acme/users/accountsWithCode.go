// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package users

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

type AccountsWithCodeAction struct {
	actionutils.ParentAction
}

func (this *AccountsWithCodeAction) RunPost(params struct {
	Code string
}) {
	accountsResp, err := this.RPC().ACMEProviderAccountRPC().FindAllACMEProviderAccountsWithProviderCode(this.AdminContext(), &pb.FindAllACMEProviderAccountsWithProviderCodeRequest{AcmeProviderCode: params.Code})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	var accountMaps = []maps.Map{}
	for _, account := range accountsResp.AcmeProviderAccounts {
		accountMaps = append(accountMaps, maps.Map{
			"id":   account.Id,
			"name": account.Name,
		})
	}
	this.Data["accounts"] = accountMaps

	this.Success()
}
