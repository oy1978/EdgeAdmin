// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package providers

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

func (this *ProviderAction) readEdgeDNS(provider *pb.DNSProvider, apiParams maps.Map) (maps.Map, error) {
	return maps.Map{}, nil
}
