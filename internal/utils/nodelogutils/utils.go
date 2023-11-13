// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build !plus

package nodelogutils

import (
	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeCommon/pkg/langs"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
)

// FindCommonTags 查找常用的标签
func FindNodeCommonTags(langCode langs.LangCode) []maps.Map {
	return []maps.Map{
		{
			"name": langs.Message(langCode, codes.Log_TagListener),
			"code": "LISTENER",
		},
		{
			"name": langs.Message(langCode, codes.Log_TagWAF),
			"code": "WAF",
		},
		{
			"name": langs.Message(langCode, codes.Log_TagAccessLog),
			"code": "ACCESS_LOG",
		},
	}
}
