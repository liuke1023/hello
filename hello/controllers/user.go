package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	//"github.com/astaxie/beego/session"
	mymodels "hello/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	myuser := mymodels.User{}
	myuser.Uname = this.GetString("uname")
	myuser.Upw = this.GetString("upw")
	r, err := myuser.Login()
	if err != nil {
		fmt.Print(err)
	}
	if r {
		sess, _ := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		defer sess.SessionRelease(this.Ctx.ResponseWriter)
		sess.Set("username", myuser.Uname)
		sess.Set("loginstatus", bool(true))
		this.Ctx.WriteString("yes")
	} else {
		this.Ctx.WriteString("No")
	}

}

func LoginCheck(w http.ResponseWriter, r *http.Request) bool {
	sess, _ := globalSessions.SessionStart(w, r)
	w.Header().Set("Content-Type", "text/html")
	v := sess.Get("loginstatus")
	if v == nil {
		return false
	}
	return v.(bool)

}
