package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// GmController operations for Gm
type GmController struct {
	beego.Controller
}

// URLMapping ...
func (c *GmController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Get ...
func (c *GmController) Get() {
	c.Data["HTTPSERVER"] = "127.0.0.1:8001/script"
	c.TplName = "gm/index.html"
}

// Post ...
// @Title Create
// @Description create Gm
// @Param	body		body 	models.Gm	true		"body for Gm content"
// @Success 201 {object} models.Gm
// @Failure 403 body is empty
// @router / [post]
func (c *GmController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Gm by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Gm
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GmController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Gm
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Gm
// @Failure 403
// @router / [get]
func (c *GmController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Gm
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Gm	true		"body for Gm content"
// @Success 200 {object} models.Gm
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GmController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Gm
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *GmController) Delete() {

}
