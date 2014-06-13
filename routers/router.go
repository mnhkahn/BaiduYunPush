package routers

import (
	"BaiduYunPush/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/push", &controllers.MainController{}, "get:Push")

}
