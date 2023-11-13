package actionutils

import (
	"net/http"

	"github.com/iwind/TeaGo/actions"
	"github.com/oy1978/EdgeAdmin/internal/csrf"
)

type CSRF struct {
}

func (this *CSRF) BeforeAction(actionPtr actions.ActionWrapper, paramName string) (goNext bool) {
	action := actionPtr.Object()
	token := action.ParamString("csrfToken")
	if !csrf.Validate(token) {
		action.ResponseWriter.WriteHeader(http.StatusForbidden)
		action.WriteString("表单已失效，请刷新页面后重试(001)")
		return
	}

	return true
}
