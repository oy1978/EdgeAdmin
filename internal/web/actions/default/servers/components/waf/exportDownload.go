package waf

import (
	"strconv"

	"github.com/iwind/TeaGo/types"
	"github.com/oy1978/EdgeAdmin/internal/ttlcache"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
)

type ExportDownloadAction struct {
	actionutils.ParentAction
}

func (this *ExportDownloadAction) Init() {
	this.Nav("", "", "")
}

func (this *ExportDownloadAction) RunGet(params struct {
	Key      string
	PolicyId int64
}) {
	item := ttlcache.DefaultCache.Read(params.Key)
	if item == nil || item.Value == nil {
		this.WriteString("找不到要导出的内容")
		return
	}

	ttlcache.DefaultCache.Delete(params.Key)

	data, ok := item.Value.([]byte)
	if ok {
		this.AddHeader("Content-Disposition", "attachment; filename=\"WAF-"+types.String(params.PolicyId)+".json\";")
		this.AddHeader("Content-Length", strconv.Itoa(len(data)))
		_, _ = this.Write(data)
	} else {
		this.WriteString("找不到要导出的内容")
		return
	}
}
