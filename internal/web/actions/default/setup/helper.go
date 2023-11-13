package setup

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/setup"
)

type Helper struct {
}

func (this *Helper) BeforeAction(actionPtr actions.ActionWrapper) (goNext bool) {
	if setup.IsConfigured() {
		actionPtr.Object().RedirectURL("/")
		return false
	}
	return true
}
