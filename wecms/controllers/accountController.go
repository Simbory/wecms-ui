package controllers

import (
	"github.com/Simbory/wemvc"
	"github.com/Simbory/wecms"
)

type AccountController struct {
	wemvc.Controller
}

func (ctrl AccountController) GetLogin() interface{} {
	return ctrl.View()
}

func (ctrl AccountController) PostLogin() interface{} {
	returnUrl := ctrl.Request().URL.Query().Get("returnUrl")
	if len(returnUrl) == 0 {
		returnUrl = "/wecms/panel"
	}
	ctrl.Session().Set("WECMS_USER", &wecms.User{
		Id: wecms.NewID(),
		Email: "simbory@sina.cn",
		UserName: "Simbory.Lu",
		Password: "",
		FullName: "Simbory Lu",
		FirstName: "Simbory",
		LastName: "Lu",
		Roles: []wecms.RoleType{wecms.Administrator},
	})
	return ctrl.Redirect(returnUrl)
}