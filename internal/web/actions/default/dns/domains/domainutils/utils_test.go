// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package domainutils

import (
	"testing"

	"github.com/iwind/TeaGo/assert"
	"github.com/oy1978/EdgeCommon/pkg/dnsconfigs"
)

func TestValidateRecordValue(t *testing.T) {
	a := assert.NewAssertion(t)

	// A
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeA, "1.2")
		a.IsFalse(ok)
	}
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeA, "1.2.3.400")
		a.IsFalse(ok)
	}
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeA, "1.2.3.4")
		a.IsTrue(ok)
	}

	// CNAME
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeCNAME, "example.com")
		a.IsTrue(ok)
	}
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeCNAME, "example.com.")
		a.IsTrue(ok)
	}
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeCNAME, "hello, world")
		a.IsFalse(ok)
	}

	// AAAA
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeAAAA, "1.2.3.4")
		a.IsFalse(ok)
	}
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeAAAA, "2001:0db8:85a3:0000:0000:8a2e:0370:7334")
		a.IsTrue(ok)
	}

	// NS
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeNS, "1.2.3.4")
		a.IsFalse(ok)
	}
	{
		_, ok := ValidateRecordValue(dnsconfigs.RecordTypeNS, "example.com")
		a.IsTrue(ok)
	}
}
