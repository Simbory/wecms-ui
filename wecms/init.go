package wecms

import (
	"github.com/Simbory/wemvc"
	ctrl "github.com/Simbory/wecms-ui/wecms/controllers"
	"github.com/Simbory/wecms-ui/wecms/filters"
)

func init() {
	area := wemvc.GetArea("wecms")
	if area == nil {
		return
	}
	area.ServeDir("content")
	area.SetPathFilter("/", filters.LoginFilter)
	area.Route("/account/<action>", ctrl.AccountController{}, "")
	area.Route("/template/<action>/<id=>", ctrl.TemplateController{}, "")
}