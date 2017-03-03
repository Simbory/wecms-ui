package controllers

import (
	"github.com/Simbory/wecms"
	"strings"
	"time"
)

type TemplateController struct {
	authController
}

func (ctrl TemplateController) GetChildren() interface{} {
	var id = wecms.ID(strings.ToLower(ctrl.RouteData()["id"]))
	children,err := ctrl.Editing().ChildTemplateEntries(id)
	if err != nil {
		return ctrl.NotFound()
	}
	ctrl.ViewData["Children"] = children
	return ctrl.View()
}

func (ctrl TemplateController) GetNew() interface{} {
	var parent = ctrl.Request().URL.Query().Get("parent")
	if len(parent) == 0 {
		parent = string(wecms.RootID)
	}
	if parent == string(wecms.RootID) {
		ctrl.ViewData["parent"] = parent
	} else {
		pEntry,err := ctrl.Editing().GetTemplateEntry(wecms.ID(parent))
		if err != nil {
			panic(err)
		}
		if pEntry == nil {
			return ctrl.NotFound()
		}
		ctrl.ViewData["parent"] = parent
	}
	return ctrl.View()
}

func (ctrl TemplateController) PostNew() interface{} {
	var entry = wecms.TemplateEntry{
		Name: ctrl.Request().Form.Get("Name"),
		Type: ctrl.Request().Form.Get("Type"),
		Container: wecms.ID(ctrl.Request().Form.Get("Container")),
		CreatedBy: ctrl.Request().Form.Get("CreatedBy"),
		UpdatedBy: ctrl.Request().Form.Get("UpdatedBy"),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := ctrl.Editing().SaveTemplateEntry(&entry)
	if err != nil {
		panic(err)
	}
	return ctrl.PlainText("OK")
}