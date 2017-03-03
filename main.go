package main

import (
	"github.com/Simbory/wemvc"
	_ "github.com/Simbory/wecms-ui/controllers"
	"github.com/Simbory/wecms"
	_ "github.com/Simbory/wecms-ui/wecms"
)

func main() {
	wemvc.ServeDir("/content/")
	wemvc.OnAppInit(func() error {
		_,conn := wemvc.Config().GetConnConfig("master")
		master,err := wecms.NewRepository("master", conn, "master", 100, 1000)
		if err != nil {
			return err
		}
		wecms.RegRepository(master)
		_,conn = wemvc.Config().GetConnConfig("core")
		core,err := wecms.NewRepository("core", conn, "core", 100, 1000)
		if err != nil {
			return err
		}
		wecms.RegRepository(core)
		_,conn = wemvc.Config().GetConnConfig("web")
		web,err := wecms.NewRepository("web", conn, "web", 100, 1000)
		if err != nil {
			return err
		}
		wecms.RegRepository(web)
		return nil
	})
	wemvc.Run(8080)
}