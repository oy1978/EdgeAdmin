// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package utils_test

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
	"github.com/oy1978/EdgeAdmin/internal/utils"
)

func TestValidateEmail(t *testing.T) {
	var a = assert.NewAssertion(t)
	a.IsTrue(utils.ValidateEmail("aaaa@gmail.com"))
	a.IsTrue(utils.ValidateEmail("a.b@gmail.com"))
	a.IsTrue(utils.ValidateEmail("a.b.c.d@gmail.com"))
	a.IsTrue(utils.ValidateEmail("aaaa@gmail.com.cn"))
	a.IsTrue(utils.ValidateEmail("hello.world.123@gmail.123.com"))
	a.IsTrue(utils.ValidateEmail("10000@qq.com"))
	a.IsFalse(utils.ValidateEmail("aaaa.@gmail.com"))
	a.IsFalse(utils.ValidateEmail("aaaa@gmail"))
	a.IsFalse(utils.ValidateEmail("aaaa@123"))
}
