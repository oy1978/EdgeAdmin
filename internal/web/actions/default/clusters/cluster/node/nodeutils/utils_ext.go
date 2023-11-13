// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .
//go:build !plus

package nodeutils

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
)

func filterMenuItems(menuItems []maps.Map, menuItem string, prefix string, query string, info *pb.FindEnabledNodeConfigInfoResponse, langCode string) []maps.Map {
	return menuItems
}
