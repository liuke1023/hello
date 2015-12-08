package controllers

import (
	"github.com/astaxie/beego"
)

type LogMainController struct {
	beego.Controller
}

func (this *LogMainController) Get() {
	if LoginCheck(this.Ctx.ResponseWriter, this.Ctx.Request) {
		this.TplNames = "logmain.tpl"
	} else {
		this.TplNames = "index.tpl"
	}

}
