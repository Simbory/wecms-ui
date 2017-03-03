package controllers

import "github.com/Simbory/wemvc"

// HomeController home controller
type HomeController struct {
	wemvc.Controller
}

// GetIndex use http GET method to visit index action
func (home HomeController) GetIndex() interface{} {
	home.ViewData["Message"] = "Welcome to WEMVC 1.0"
	return home.View()
}

// GetAbout use http GET method to visit about action
func (home HomeController) GetAbout() interface{} {
	home.ViewData["Message"] = "About wemvc"
	return home.View()
}

// GetAbout use http GET method to visit about action
func (home HomeController) GetContact() interface{} {
	home.ViewData["Message"] = "Contact us"
	return home.View()
}