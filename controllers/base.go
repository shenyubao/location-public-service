package controllers

import "github.com/astaxie/beego"


type BaseController struct {
	beego.Controller
}

const (
	CODE_SUCCESS    = 100000
	CODE_ERR_PARAM  = 100001
	CODE_ERR_SYSTEM = 100002
	CODE_ERR_NODATA = 100003
)
