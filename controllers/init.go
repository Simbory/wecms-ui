package controllers

import "github.com/Simbory/wemvc"

func init() {
	wemvc.Route("/<action=index>", HomeController{}, "")
}