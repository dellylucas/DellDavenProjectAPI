package controllers

import (
	"DellDaven_API/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type FrontTextController struct {
	beego.Controller
}

// @Title GetTextFront
// @Description Logs user into the system
// @Success 200 {string} text
// @Failure 403 user not exist
// @router / [get]
func (u *FrontTextController) GetTextFront() {
	text, err := models.GetText("main")
	if err == nil {
		if text.Id != 0 {
			u.Data["json"] = text
		} else {
			u.Data["json"] = nil
		}
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}
