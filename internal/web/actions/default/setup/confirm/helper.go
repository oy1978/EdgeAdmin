package confirm

import (
	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/setup"
)

type Helper struct {
}

func (this *Helper) BeforeAction(actionPtr actions.ActionWrapper) (goNext bool) {
	if !setup.IsNewInstalled() {
		actionPtr.Object().RedirectURL("/")
		return false
	}
	return true
}
