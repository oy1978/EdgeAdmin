package actionutils

import (
	"context"

	"github.com/iwind/TeaGo/maps"
	"github.com/oy1978/EdgeAdmin/internal/rpc"
)

type ActionInterface interface {
	RPC() *rpc.RPCClient

	AdminContext() context.Context

	ViewData() maps.Map
}
