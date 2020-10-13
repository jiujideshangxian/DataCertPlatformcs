package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//路由router
    beego.Router("/", &controllers.MainController{})
    //用户注册接口
    beego.Router("/register", &controllers.MainController{})
}
