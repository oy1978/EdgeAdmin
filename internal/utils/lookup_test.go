// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package utils_test

import (
	"testing"

	"github.com/oy1978/EdgeAdmin/internal/utils"
)

func TestLookupCNAME(t *testing.T) {
	t.Log(utils.LookupCNAME("www.yun4s.cn"))
}
