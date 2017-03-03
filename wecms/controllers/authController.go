package controllers

import (
	"github.com/Simbory/wemvc"
	"github.com/Simbory/wecms"
	"errors"
)

type authController struct {
	User *wecms.User
	Rep  *wecms.Repository

	wemvc.Controller
}

func (ctrl *authController) Editing() *wecms.RepEditing {
	editing := ctrl.Rep.Editing(ctrl.User)
	if editing == nil {
		panic(errors.New("You do not have permission to access the repository editing."))
	}
	return editing
}

func (ctrl *authController) OnInit(ctx *wemvc.Context) {
	ctrl.Controller.OnInit(ctx)
	ctrl.Rep = wecms.GetRepository("master")
	if ctrl.Rep == nil {
		panic("The master repository cannot be found")
	}
	userObj := ctx.CtxItems().Get("WECMS_USER")
	if userObj == nil {
		panic(errors.New("Unauthorized user has no permission to access this page."))
	}
	if user,ok := userObj.(*wecms.User); ok {
		ctrl.User = user
	} else {
		panic(errors.New("Unauthorized user has no permission to access this page."))
	}
}