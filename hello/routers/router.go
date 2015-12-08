package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/yzm", &controllers.YzmController{}, "Get:GetYzm")             //验证码控制器
	beego.Router("/yzm/judgeyzm", &controllers.YzmController{}, "Post:JudgeYzm") //验证码ajax验证
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/logmain", &controllers.LogMainController{})
}
