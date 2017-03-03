package filters

import (
	"github.com/Simbory/wemvc"
	"net/http"
)

func LoginFilter(ctx *wemvc.Context) {
	if ctx.Ctrl.ControllerName != "account" {
		session := ctx.Session()
		user := session.Get("WECMS_USER")
		if user != nil {
			ctx.CtxItems().Set("WECMS_USER", user)
			return
		}
		http.Redirect(ctx.Response(), ctx.Request(), "/wecms/account/login?returnUrl=" + ctx.Request().URL.String(), 302)
		ctx.EndContext()
	}
}
