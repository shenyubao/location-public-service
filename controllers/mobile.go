package controllers

import (
	"regexp"
	"souche-public-service/models"
	"fmt"
)

// Operations about Mobile
type MobileController struct {
	BaseController
}

func (u *MobileController) GetLocation() {
	_mobile := u.GetString("mobile")
	if m, _ := regexp.MatchString(`^(1[3|4|5|8|7][0-9]\d{4,8})$`, _mobile); !m {
		u.Data["json"] = models.Render(CODE_ERR_PARAM, "error", "")
		u.ServeJson()
	}else{
		fmt.Println(_mobile)
		area := models.FindLocation(_mobile[0:7])
		u.Data["json"] = models.Render(CODE_SUCCESS, "success", area)
		u.ServeJson()
	}
}
