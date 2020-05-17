package controllers

import (
	"DellDaven_API/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.Members	true		"body for user content"
// @Success 200 {int} models.Members.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.Members
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := models.AddUser(&user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = true
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		/*	user, err := models.GetUser(uid)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = user
			}*/
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.Members
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		/*uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}*/
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	//uid := u.GetString(":uid")
	//models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	models.Members	true		"body for user content"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var user models.Members
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := models.Login(&user)
	if err == nil {
		if user.Id != 0 {
			u.Data["json"] = true
		} else {
			u.Data["json"] = false
		}
	} else {
		u.Data["json"] = false
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
