// Copyright 2023 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package apinodeutils_test

import (
	"runtime"
	"testing"

	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/oy1978/EdgeAdmin/internal/utils/apinodeutils"
)

func TestUpgrader_CanUpgrade(t *testing.T) {
	t.Log(apinodeutils.CanUpgrade("0.6.3", runtime.GOOS, runtime.GOARCH))
}

func TestUpgrader_Upgrade(t *testing.T) {
	var upgrader = apinodeutils.NewUpgrader(1)
	err := upgrader.Upgrade()
	if err != nil {
		t.Fatal(err)
	}
}
