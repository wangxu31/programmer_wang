package routers

import (
	"programmer_wang/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/wx", &controllers.WeChatController{}, "get:Verify;post:Answer")
}
