package recovers

import (
	"github.com/iwind/TeaGo/actions"
	teaconst "github.com/oy1978/EdgeAdmin/internal/const"
)

type Helper struct {
}

func (this *Helper) BeforeAction(actionPtr actions.ActionWrapper) (goNext bool) {
	if !teaconst.IsRecoverMode {
		actionPtr.Object().RedirectURL("/")
		return false
	}
	return true
}
