package controllers

import (
	"beegoWeb/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

// PlayerController operations for Player
type PlayerController struct {
	beego.Controller
}

// URLMapping ...
func (c *PlayerController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *PlayerController) Get() {
	// 使用 GetString 方法获取查询参数
	data := c.GetString("message")
	//data1 := strings.Replace(data, "=", " ", -1)

	var player models.Player

	// 解析JSON数据
	err := json.Unmarshal([]byte(data), &player)
	if err != nil {
		fmt.Println("解析JSON数据出错:", err)
		return
	}

	// 打印解析后的结果
	fmt.Println("玩家id:", player.Id)
	fmt.Printf("玩家 data: %+v\n", player)

	c.Data["AccId"] = player.AccId
	c.Data["Id"] = player.Id
	c.Data["Name"] = player.Name
	c.Data["Coin"] = player.TotalCoin
	c.Data["Diamond"] = player.TotalDiamond
	c.Data["Power"] = player.Power
	c.Data["VipLevel"] = player.VipLevel
	c.Data["LastOnlineTime"] = player.LastLoginTime
	c.Data["MainBuildingLevel"] = player.MainBuildingLevel
	c.Data["AllianceId"] = player.AllianceId
	c.TplName = "player/index.html"
}

// Post ...
// @Title Create
// @Description create Player
// @Param	body		body 	models.Player	true		"body for Player content"
// @Success 201 {object} models.Player
// @Failure 403 body is empty
// @router / [post]
func (c *PlayerController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Player by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Player
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PlayerController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Player
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Player
// @Failure 403
// @router / [get]
func (c *PlayerController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Player
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Player	true		"body for Player content"
// @Success 200 {object} models.Player
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PlayerController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Player
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PlayerController) Delete() {

}
