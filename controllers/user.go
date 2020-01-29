package controllers

import (
	"encoding/json"
	"golemonilo/models"

	lemonilo "github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	lemonilo.Controller
}

type Response struct {
	Errcode int         `json:"code"`
	Errmsg  string      `json:"message"`
	Data    interface{} `json:"data"`
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = Response{Errmsg: "200", Errcode: 200, Data: uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = Response{Errmsg: "200", Errcode: 200, Data: users}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid != 0 {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = Response{Errmsg: "500", Errcode: 500, Data: err.Error()}
		} else {
			u.Data["json"] = Response{Errmsg: "success", Errcode: 200, Data: user}
		}
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
	uid, _ := u.GetInt(":uid")
	if uid != 0 {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		user.ID = uid
		uu, err := models.UpdateUser(user)
		if err != nil {
			u.Data["json"] = Response{Errmsg: "500", Errcode: 500, Data: err.Error()}
		} else {
			u.Data["json"] = Response{Errmsg: "success", Errcode: 200, Data: uu}
		}
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
	uid, _ := u.GetInt(":uid")
	if models.DeleteUser(uid) != 0 {
		u.Data["json"] = Response{Errmsg: "200", Errcode: 200, Data: "Success"}
		u.ServeJSON()
		return
	}
	u.Data["json"] = Response{Errmsg: "Failed to delete data", Errcode: 400, Data: ""}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	email			query 	string	true		"The email for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var users models.User

	session := u.StartSession()
	auth := session.Get("token")

	if auth != nil {
		u.Data["json"] = Response{Errmsg: "logged in", Errcode: 200, Data: ""}
	} else {
		json.Unmarshal(u.Ctx.Input.RequestBody, &users)
		isMatch, user := models.Login(users)
		if isMatch {
			session.Set("token", user.Token)
			u.Data["json"] = Response{Errmsg: "200", Errcode: 200, Data: user}
		} else {
			u.Data["json"] = Response{Errmsg: "user not exist / wrong password", Errcode: 400, Data: ""}
		}
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.DestroySession()
	u.Data["json"] = Response{Errmsg: "success logout", Errcode: 200, Data: ""}
	u.ServeJSON()
}
