package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}


/**
 * 直接跳转展示用户登录页面
 */

func (l *LoginController)Get()  {
	l.TplName="login.html"
}


/**
 * post方法处理用户的登录请求
 */
func (l *LoginController)Post()  {
	var user = 
}