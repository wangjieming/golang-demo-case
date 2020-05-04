package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"go-demo/models"
)

// Operations about Users
type BlogController struct {
	beego.Controller
}



// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:id [get]
func (u *BlogController) GetOneBlog()  {
	//id := u.GetString(":id")
	id,err := u.GetInt(":id")
	if err == nil {
		blog, err := models.GetBlogText(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = blog
		}

		bytes, err := json.MarshalIndent(blog, "", " ")
		fmt.Println("------- blog -------" + string(bytes))
	}
	u.ServeJSON()
}
