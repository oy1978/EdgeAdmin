// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build !plus

package locationutils

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs"
)

func (this *LocationHelper) filterMenuItems1(locationConfig *serverconfigs.HTTPLocationConfig, menuItems []maps.Map, serverIdString string, locationIdString string, secondMenuItem string, actionPtr actions.ActionWrapper) []maps.Map {
	return menuItems
}

func (this *LocationHelper) filterMenuItems2(locationConfig *serverconfigs.HTTPLocationConfig, menuItems []maps.Map, serverIdString string, locationIdString string, secondMenuItem string, actionPtr actions.ActionWrapper) []maps.Map {
	return menuItems
}

func (this *LocationHelper) filterMenuItems3(locationConfig *serverconfigs.HTTPLocationConfig, menuItems []maps.Map, serverIdString string, locationIdString string, secondMenuItem string, actionPtr actions.ActionWrapper) []maps.Map {
	return menuItems
}
