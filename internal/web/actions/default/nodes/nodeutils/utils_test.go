package nodeutils

import (
	"testing"

	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/logs"
	"github.com/oy1978/EdgeAdmin/internal/rpc"
)

func TestSendMessageToCluster(t *testing.T) {
	rpcClient, err := rpc.SharedRPC()
	if err != nil {
		t.Fatal(err)
	}
	ctx := rpcClient.Context(1)

	results, err := SendMessageToCluster(ctx, 1, "test", nil, 30)
	if err != nil {
		t.Fatal(err)
	}
	logs.PrintAsJSON(results, t)
}
