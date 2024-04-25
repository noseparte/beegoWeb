package controllers

import (
	"beegoWeb/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Get ...
func (c *UserController) Get() {
	// 使用 GetString 方法获取查询参数
	//data := c.GetString("message")
	//data1 := strings.Replace(data, "=", " ", -1)

	// 假设我们有以下JSON数据
	jsonData := `{
		"id": "918513",
		"x": 207,
		"y": 7,
		"allianceId": "4333622005121",
		"hp": 0,
		"fireTime": "0",
		"lastTime": "0",
		"status": 2
	}`

	// 创建一个User类型的变量
	var user models.PlayerCity

	// 解析JSON数据
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("解析JSON数据出错:", err)
		return
	}

	// 打印解析后的结果
	fmt.Println("玩家id:", user.ID)
	fmt.Printf("玩家 data: %+v\n", user)

	//c.Data["user"] = user
	//c.TplName = "user/index.html"
}

// Post ...
// @Title Create
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {object} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {

}
